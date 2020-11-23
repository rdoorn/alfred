package zwave

import (
	"fmt"
	"log"
)

type Frame interface {
	Head() byte
	Length() int
	Type() ZWaveType
	CommandID() byte
	CommandParams() []byte
	Checksum() byte
	ToBytes() []byte
	String() string
}

type ZwaveFrame struct {
	head             byte
	length           int
	frametype        ZWaveType // REQ: 0x00 RES: 0x01
	apiCommand       byte
	apiCommandParams []byte
	checksum         byte
}

const (
	FrameSOF = 0x01
	FrameACK = 0x06
	FrameNAK = 0x15
	FrameCAN = 0x18

	REQ = 0x00
	RES = 0x01
)

func ACK() Frame {
	return &ZwaveFrame{head: FrameACK}
}

func NAK() Frame {
	return &ZwaveFrame{head: FrameNAK}
}

func CAN() Frame {
	return &ZwaveFrame{head: FrameCAN}
}

func SOF(ftype ZWaveType, command byte, param ...byte) Frame {
	z := &ZwaveFrame{head: FrameSOF, frametype: ftype, apiCommand: command, apiCommandParams: param}
	return z.Pack()
}

func (f *ZwaveFrame) Pack() *ZwaveFrame {
	if f.head == FrameSOF {
		f.length = len(f.apiCommandParams) + 3
		f.checksum = f.CalculateChecksum()
	}
	return f
}

func (f *ZwaveFrame) CalculateChecksum() byte {
	data := []byte{0xFF, byte(f.length), byte(f.frametype), f.apiCommand}
	data = append(data, f.apiCommandParams...)

	var offset int
	ret := data[offset]
	for i := offset + 1; i < len(data); i++ {
		// Xor bytes
		log.Printf("XOR ret %.8b data %.8b (%d) ", ret, data[i], data[i])
		ret ^= data[i]
	}
	return ret
}

func (f *ZwaveFrame) Head() byte            { return f.head }
func (f *ZwaveFrame) Length() int           { return f.length }
func (f *ZwaveFrame) Type() ZWaveType       { return f.frametype }
func (f *ZwaveFrame) CommandID() byte       { return f.apiCommand }
func (f *ZwaveFrame) CommandParams() []byte { return f.apiCommandParams }
func (f *ZwaveFrame) Checksum() byte {
	if f.checksum == 0x00 {
		return f.CalculateChecksum()
	}
	return f.checksum
}

func (f *ZwaveFrame) ToBytes() []byte {
	switch f.Head() {
	case FrameACK, FrameCAN, FrameNAK:
		return []byte{f.Head()}
	default:
		data := []byte{f.Head(), byte(f.Length()), byte(f.Type()), f.CommandID()}
		data = append(data, f.CommandParams()...)
		data = append(data, f.Checksum())
		return data
	}
}

func NewFrame(data []byte) (int, Frame, error) {
	switch data[0] {
	case FrameACK:
		log.Printf("Attempt to create frame based on: %x", data[0])
		return 1, ACK(), nil
	case FrameNAK:
		log.Printf("Attempt to create frame based on: %x", data[0])
		return 1, NAK(), nil
	case FrameCAN:
		log.Printf("Attempt to create frame based on: %x", data[0])
		return 1, CAN(), nil
	case FrameSOF:
		if len(data) < 2 {
			//log.Printf("not all data is in yet for SOF Len len=%d want=1", len(data))

			// not enough data yet, we need atleast 2 characters to find the len
			return 2, nil, nil
		}

		length := int(data[1])
		if len(data) < length+2 { // include SOF+Len+checksum -1 for counting from 0
			//log.Printf("not all data is in yet for SOF Body len=%d want=%d", len(data), length+2)
			// not enough data yet, we need atleast 2 characters + length of message + checksum
			return length + 2, nil, nil
		}

		/* FIXME:

		    2020/11/18 17:23:50 Attempt to create frame based on: "\x01\x02\x05"
		panic: runtime error: slice bounds out of range [4:3]

		goroutine 6 [running]:
		github.com/rdoorn/alfred/pkg/zwave.NewFrame(0xc00001c282, 0x25, 0x3e, 0x1, 0x1, 0x0, 0x0, 0x0)
			/Users/rdoorn/Work/go/src/github.com/rdoorn/alfred/pkg/zwave/frame.go:128 +0x62f
		github.com/rdoorn/alfred/pkg/zwave.(*ZwaveAPI).Reader(0xc00000e0c0)
			/Users/rdoorn/Work/go/src/github.com/rdoorn/alfred/pkg/zwave/transport.go:57 +0x97
		created by github.com/rdoorn/alfred/pkg/zwave.New
			/Users/rdoorn/Work/go/src/github.com/rdoorn/alfred/pkg/zwave/controller.go:47 +0x494
		exit status 2

		*/
		// put data in SOF to create frame
		log.Printf("Attempt to create frame based on: %x", data[:length+1])
		frame := SOF(ZWaveType(data[2]), data[3], data[4:length+1]...)
		if frame.Checksum() != data[length+1] {
			return -1, nil, fmt.Errorf("Checksum mismatch: (%d vs %d) %+v", frame.Checksum(), data[length+1], frame)
		}
		return length + 2, frame, nil
	default:
		// not a valid start char
		return -1, nil, nil
	}
}

func (f *ZwaveFrame) String() string {
	switch f.head {
	case FrameACK:
		return "ACK"
	case FrameNAK:
		return "NAK"
	case FrameCAN:
		return "CAN"
	case FrameSOF:
		return fmt.Sprintf("SOF length:%d %s, Command: %q, Params: %q, Checksum: %q", f.length, f.Type(), f.apiCommand, f.apiCommandParams, f.checksum)
	}
	return fmt.Sprintf("Unknown frame; %+v", f.head)
}
