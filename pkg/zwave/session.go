package zwave

import (
	"log"
)

type Session interface {
	//WriteFunction(f ZWaveFunction)
	Write(Frame) error
	Read() Frame
	Reader()
}

type ZwaveSession struct {
	t Transport
	//in  <-chan Message
	//out chan<- Message
}

//func (z *ZwaveSession) In() <-chan Message {
//return z.in
//}

//func (z *ZwaveSession) Out() chan<- Message {
//return z.out
//}

func NewZwaveSession(t Transport) *ZwaveSession {
	return &ZwaveSession{
		t: t,
		//in:  make(<-chan Message),
		//out: make(chan<- Message),
	}
}

func (z *ZwaveSession) Write(f Frame) error {
	w := z.t.Write()
	log.Printf("Session write: %+v", f)
	w <- f
	return nil
}

/*func (z *ZwaveSession) WriteFunction(f ZWaveFunction) error {
	return z.t.Writer(f)
}*/

func (z *ZwaveSession) Read() Frame {
	f := <-z.t.Read()
	log.Printf("Session read: %+v", f)
	if f.Type() == RES {
		z.Write(ACK())
	}
	return f
}

func (z *ZwaveSession) Reader() {
	for {
		select {
		case f := <-z.t.Read():
			log.Printf("Session read: %+v", f)

			// send a ACK to any received response
			if f.Type() == RES {
				z.Write(ACK())
			}

			switch f.CommandID() {
			case DiscoveryNodes:
				log.Printf("We have a discovered nodes report")
				log.Printf("ver: %d", f.CommandParams()[0])
				log.Printf("cap: %d", f.CommandParams()[1])
				log.Printf("29: %d", f.CommandParams()[2])
				for index, bits := range f.CommandParams()[3 : len(f.CommandParams())-2] {
					log.Printf("bit: %d", bits)
					log.Printf("node: %d state: %t", index*8+0, bits&0x01 != 0)
					log.Printf("node: %d state: %t", index*8+1, bits&0x02 != 0)
					log.Printf("node: %d state: %t", index*8+2, bits&0x04 != 0)
					log.Printf("node: %d state: %t", index*8+3, bits&0x08 != 0)
					log.Printf("node: %d state: %t", index*8+4, bits&0x10 != 0)
					log.Printf("node: %d state: %t", index*8+5, bits&0x20 != 0)
					log.Printf("node: %d state: %t", index*8+6, bits&0x40 != 0)
					log.Printf("node: %d state: %t", index*8+7, bits&0x80 != 0)
				}
				log.Printf("chiptype: %d", f.CommandParams()[len(f.CommandParams())-2])
				log.Printf("chipversion: %d", f.CommandParams()[len(f.CommandParams())-1])

				// TEMP:
				z.Write(SOF(REQ, GetNodeProtocolInfo, 3))
			case GetNodeProtocolInfo:
				log.Printf("Node protocol information received")
				log.Printf("Listening (awake): %t", f.CommandParams()[0]&0x80 != 0)
				log.Printf("Routing: %t", f.CommandParams()[0]&0x40 != 0)
				log.Printf("Version: %d", f.CommandParams()[0]&0x07+1)

				baud := 9600
				if f.CommandParams()[0]&0x38 == 0x10 {
					baud = 40000
				}
				log.Printf("Max Baud: %d", baud)

				log.Printf("Flirs: %t", f.CommandParams()[1]&0x60 != 0)
				log.Printf("Beaming: %t", f.CommandParams()[1]&0x10 != 0)
				log.Printf("Security: %t", f.CommandParams()[1]&0x01 != 0)

				log.Printf("Reserved: %d", f.CommandParams()[2])
				log.Printf("Basic: %d", f.CommandParams()[3])
				log.Printf("Generic: %d", f.CommandParams()[4])
				log.Printf("Specific: %d", f.CommandParams()[5])
				z.Write(SOF(REQ, GetControllerCapabilities))
			case GetControllerCapabilities:
				log.Printf("Controller is secondary: %t", f.CommandParams()[0]&0x01 != 0)
				log.Printf("Controller on other network: %t", f.CommandParams()[0]&0x02 != 0)
				log.Printf("Controller nodeid server present: %t", f.CommandParams()[0]&0x04 != 0)
				log.Printf("Controller is real primary: %t", f.CommandParams()[0]&0x08 != 0)
				log.Printf("Controller is suc: %t", f.CommandParams()[0]&0x10 != 0)
				log.Printf("Controller no nodes included: %t", f.CommandParams()[0]&0x20 != 0)
				//z.Write(SOF(REQ, SendData, 0x03, 0x04))
				// request spec // node id

				data := []byte{0x04, 0x00}
				l := len(data)
				// node man spec info
				data = append(data, 0x01) // ack
				data = append(data, 0x50) // callback
				data = append([]byte{0x03, byte(l)}, data...)
				z.Write(SOF(REQ, SendData, data...))
				// node, len, data, tx

				// REQ, senddata, node, lenght
				// nodeid , pdata, datalen, txoptions, txstatus
			case 0x13:
				z.Write(ACK())
				log.Printf("send data req to software")
				log.Printf("len %d", f.CommandParams()[0])
				//log.Printf("id %d", f.CommandParams()[2])

			}
		}
	}
}
