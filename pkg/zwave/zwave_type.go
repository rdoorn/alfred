package zwave

type ZWaveType byte

const (
	Request  = 0x00
	Response = 0x01
)

func (b ZWaveType) String() string {
	str := ""
	switch b {
	case 0x00:
		str = "Request"
	case 0x01:
		str = "Response"
	default:
		str = "Unknown"
	}

	return str
}
