package zwave

// ZWaveBasicType contains the basic type of the protocol information
type ZWaveBasicType byte

// Protocol information: basic type
const (
	BasicTypeController       = 0x01 /*Node is a portable controller */
	BasicTypeStaticController = 0x02 /*Node is a static controller*/
	BasicTypeSlave            = 0x03 /*Node is a slave*/
	BasicTypeRoutingSlave     = 0x04 /*Node is a slave with routing capabilities*/
)

func (b ZWaveBasicType) String() string {
	str := ""
	switch b {
	case 0x01:
		str = "BasicTypeController"
	case 0x02:
		str = "BasicTypeStaticController"
	case 0x03:
		str = "BasicTypeSlave"
	case 0x04:
		str = "BasicTypeRoutingSlave"
	default:
		str = "Unknown"
	}

	return str
}
