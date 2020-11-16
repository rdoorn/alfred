package zwave

type Session interface {
	Write(Message) error
	Read() Message
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

func (z *ZwaveSession) Write(m Message) error {
	return z.t.Writer(m)
}

func (z *ZwaveSession) Read() Message {
	return z.t.Reader()
}
