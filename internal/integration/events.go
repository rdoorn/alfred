package integration

// Events
const (
	EventAlarmState   = iota
	EventSensorMotion // sensor readings
	EventSensorTemperature
	EventSensorHumidity
	EventSensorCo2
	EventSensorNoise
	EventSensorLux
	EventSwitchState // switch/dimmer readings
	EventDimmerLevel
	EventUsageKwh // power readings
	EventUsageWatt
)

// Event is the type of event that has occured
type Event struct {
	Type    int
	Version string
	Data    interface{}
}

type SensorLux struct {
	ID  int
	Lux int
}

func (e *Event) Lux() int {
	if e.Type == EventSensorLux {
		return e.Data.(SensorLux).Lux
	}
	return 0
}

func (e *Event) ID() int {
	if e.Type == EventSensorLux {
		return e.Data.(SensorLux).ID
	}
	return 0
}
