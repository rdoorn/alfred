package zwave

// ZWaveBasicType contains the basic type of the protocol information
type ZWaveGenericTypeGenericController byte

// Protocol information: basic type
const (
	GENERIC_TYPE_GENERIC_CONTROLLER = 0x01 /*Remote Controller*/
)

func (b ZWaveGenericTypeGenericController) String() string {
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
