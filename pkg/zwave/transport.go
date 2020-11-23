package zwave

import (
	"fmt"
	"io"
	"log"
)

type Transport interface {
	Reader()
	Writer()
	Read() <-chan Frame
	Write() chan<- Frame
	Hook(rw io.ReadWriteCloser) error
	String() string
}

type ZwaveAPI struct {
	rw    io.ReadWriteCloser
	read  chan Frame
	write chan Frame
}

func NewZwaveAPI() *ZwaveAPI {
	return &ZwaveAPI{
		read:  make(chan Frame),
		write: make(chan Frame),
	}
}

func (a *ZwaveAPI) Hook(rw io.ReadWriteCloser) error {
	if rw == nil {
		return fmt.Errorf("empty hook provided in Transporter")
	}
	a.rw = rw
	return nil
}

func (a *ZwaveAPI) Reader() {
	incomming := make([]byte, 0)

	for {
		buf := make([]byte, 128)

		n, err := a.rw.Read(buf)
		if err != nil {
			log.Printf("transport read failed: %s", err)
			a.rw.Close()
			return
		}

		//log.Printf("transport read: %q (len:%d)", buf, n)

		incomming = append(incomming, buf[:n]...)

		for len(incomming) > 0 {
			length, frame, err := NewFrame(incomming)
			if err != nil {
				log.Printf("failed to create frame from incomming data: %s", err)
			}

			if length > len(incomming) { // not all data is in yet
				break
			}

			if length == -1 { // an error occured
				incomming = incomming[1:] // remove first char to try again
				log.Printf("removing first byte from buffer len=%d", len(incomming))
				continue
			}

			if length == 1 {
				// ack nak can
				log.Printf("Recieved: %s", frame)
				a.read <- frame
				incomming = incomming[1:]
				continue
			}

			incomming = incomming[length:]
			if frame != nil {
				log.Printf("Transport is sending message to bus: %+v", frame)
				a.read <- frame
			}
		} // for len incomming
	} // for
}

func (a *ZwaveAPI) Writer() {
	for {
		select {
		case frame := <-a.write:
			log.Printf("Writer received Frame: %+v", frame)
			log.Printf("Writing: %x", frame.ToBytes())
			b, err := a.rw.Write(frame.ToBytes())
			if err != nil {
				log.Printf("Error writing %d bytes: %s", b, err)

			}
			//log.Printf("Wrote %d bytes", b)
		}
	}
}

func (a *ZwaveAPI) Read() <-chan Frame {
	return a.read
}

func (a *ZwaveAPI) Write() chan<- Frame {
	return a.write
}

func (a *ZwaveAPI) String() string {
	return "ZwaveAPI"
}
