package zwave

import (
	"fmt"
	"io"
	"log"
)

type Transport interface {
	Reader()
	Writer()
	Read() <-chan Message
	Write() chan<- Message
	Hook(rw io.ReadWriteCloser) error
	String() string
}

type ZwaveAPI struct {
	rw    io.ReadWriteCloser
	read  <-chan Message
	write chan<- Message
}

func NewZwaveAPI() *ZwaveAPI {
	return &ZwaveAPI{
		read:  make(<-chan Message),
		write: make(chan<- Message),
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
	for {
		buf := make([]byte, 128)

		n, err := a.rw.Read(buf)
		if err != nil {
			log.Printf("transport read failed: ", err)
			a.rw.Close()
		}

		log.Printf("transport read: %q", buf)

	}
}

func (a *ZwaveAPI) Writer() {

}

func (a *ZwaveAPI) Read() <-chan Message {
	return a.read
}

func (a *ZwaveAPI) Write() chan<- Message {
	return a.write
}

func (a *ZwaveAPI) String() string {
	return "ZwaveAPI"
}
