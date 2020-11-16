package zwave

import (
	"fmt"
	"io"

	"github.com/tarm/serial"
)

type Device interface {
	Open() (io.ReadWriteCloser, error)
	String() string
}

type Serial struct {
	Name string
	Baud int
}

func NewSerial(dev string, rate int) *Serial {
	return &Serial{
		Name: dev,
		Baud: rate,
	}
}

func (d *Serial) Open() (io.ReadWriteCloser, error) {
	c := &serial.Config{Name: d.Name, Baud: d.Baud}
	return serial.OpenPort(c)
}

func (d *Serial) String() string {
	return fmt.Sprintf("USB Serial device %s with baudrate %d", d.Name, d.Baud)
}
