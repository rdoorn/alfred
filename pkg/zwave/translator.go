package zwave

type Translator interface {
	ByteToMessage() Message
}

type ZwaveTranslator struct{}

func NewTranslator() *ZwaveTranslator {
	return &ZwaveTranslator{}
}

func ByteToMessage(b []byte) Message {

}

func MessageToByte(m Message) []byte {

}
