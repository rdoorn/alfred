package zwave

import "fmt"

// Frame identifiers
const (
	FrameACK  = 0x06
	FrameNAK  = 0x15
	FrameCAN  = 0x18
	FrameData = 0x01
)

// FrameData Details
const (
	FrameDataREQ = 0x00
	FrameDataRES = 0x01
)

type Message interface {
	FrameType() byte
	Length() byte
	Request() byte
	Response() byte
	ApiCommand() byte
	ApiCommandParams() []byte
	Checksum() []byte
}

type ZwaveMessage struct {
	SOF              byte
	Length           byte
	Type             byte // REQ: 0x00 RES: 0x01
	ApiCommand       byte
	ApiCommandParams []byte
	Checksum         byte
}

func NewMessage(req byte, command byte, params ...byte) *ZwaveMessage {
	f := &ZwaveMessage{
		SOF:              0x01,
		Length:           byte(len(params) + 3),
		Type:             req,
		ApiCommand:       command,
		ApiCommandParams: params,
	}
	f.Checksum = f.CalculateChecksum()
	return f
}

func (f *ZwaveMessage) CalculateChecksum() byte {
	data := []byte{0x00, f.Length, f.Type, f.ApiCommand, 0x00}
	data = append(data, f.ApiCommandParams...)

	var offset int
	ret := data[offset]
	for i := offset + 1; i < len(data)-1; i++ {
		// Xor bytes
		ret ^= data[i]
	}
	// Not result
	ret = (byte)(^ret)
	return ret
}

func (f *ZwaveMessage) Packet() []byte {
	data := []byte{f.SOF, f.Length, f.Type, f.ApiCommand}
	data = append(data, f.ApiCommandParams...)
	data = append(data, f.Checksum)
	return data
}

func Decode(data []byte) (length int, msg *ZwaveMessage, err error) {
	return -1, nil, fmt.Errorf("error")
}
