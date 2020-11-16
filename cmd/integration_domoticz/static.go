package main

// alfred/livingroom/sensor/+/lux

type ZwaveDeviceDetail struct {
	Location string
	Name     string
	Type     int
}

const (
	ZwaveController = iota
	FibaroDimmer2
	FibaroRollerShutter2
	FibaroMotionSensor5
	FibaroDoubleRelay
	FibaroWallplug5
)

var (
	ZwaveDeviceDetails = map[string]ZwaveDeviceDetail{
		"01": ZwaveDeviceDetail{"Meterkast", "Z-Stick Gen5", ZwaveController},
		"03": ZwaveDeviceDetail{"Woonkamer", "Lamp Groen", FibaroDimmer2},
		"04": ZwaveDeviceDetail{"Woonkamer", "Lamp Hal", FibaroDimmer2},
		"05": ZwaveDeviceDetail{"Eetkamer", "Lamp", FibaroDimmer2},
		"06": ZwaveDeviceDetail{"Woonkamer", "Lamp Tuin", FibaroDimmer2},
		"07": ZwaveDeviceDetail{"Keuken", "Spots", FibaroDimmer2},
		"08": ZwaveDeviceDetail{"Keuken", "Lamp", FibaroDimmer2},
		"09": ZwaveDeviceDetail{"Sanne", "Lamp", FibaroDimmer2},
		"10": ZwaveDeviceDetail{"Lotte", "Lamp", FibaroDimmer2},
		"11": ZwaveDeviceDetail{"Badkamer", "Lamp", FibaroDimmer2},
		"12": ZwaveDeviceDetail{"Badkamer", "Spots", FibaroDimmer2},
		"13": ZwaveDeviceDetail{"Computerkamer", "Lamp", FibaroDimmer2},
		// DEAD/REMOVED "14": ZwaveDeviceDetail{"Meterkast", "Z-Stick Gen5"},
		"15": ZwaveDeviceDetail{"Woonkamer", "Zonnescherm", FibaroRollerShutter2},
		"16": ZwaveDeviceDetail{"Buiten", "Verlichting Tuin", FibaroDimmer2},
		"17": ZwaveDeviceDetail{"Eetkamer", "Spots", FibaroDimmer2},
		"18": ZwaveDeviceDetail{"Keuken", "Sensor", FibaroMotionSensor5},
		"19": ZwaveDeviceDetail{"Slaapkamer", "Lamp", FibaroDimmer2},
		"20": ZwaveDeviceDetail{"Buiten", "Lamp Voordeur", FibaroDoubleRelay},
		"21": ZwaveDeviceDetail{"Badkamer", "Sensor 1", FibaroMotionSensor5},
		"22": ZwaveDeviceDetail{"Badkamer", "Sensor 2", FibaroMotionSensor5},
		"23": ZwaveDeviceDetail{"Woonkamer", "Sensor 1", FibaroMotionSensor5},
		"24": ZwaveDeviceDetail{"Woonkamer", "Sensor 2", FibaroMotionSensor5},
		"25": ZwaveDeviceDetail{"Washok", "Stekker", FibaroWallplug5},
		"26": ZwaveDeviceDetail{"Washok", "Sensor", FibaroMotionSensor5},
		"27": ZwaveDeviceDetail{"Keuken", "Koelkast Stekker", FibaroWallplug5},
		"28": ZwaveDeviceDetail{"Washok", "Vriezer Stekker", FibaroWallplug5},
		"29": ZwaveDeviceDetail{"Woonkamer", "Tv Meubel Stekker", FibaroWallplug5},
		"30": ZwaveDeviceDetail{"Woonkamer", "Ziggo Stekker", FibaroWallplug5},
		"31": ZwaveDeviceDetail{"Meterkast", "Lage Stekker", FibaroWallplug5},
		"32": ZwaveDeviceDetail{"Meterkast", "Hoge Stekker", FibaroWallplug5},
		"33": ZwaveDeviceDetail{"Computerkamer", "Stekker", FibaroWallplug5},
		"34": ZwaveDeviceDetail{"Variabel", "Misc Stekker", FibaroWallplug5},
		"35": ZwaveDeviceDetail{"Variabel", "Misc2 Stekker", FibaroWallplug5},
		"36": ZwaveDeviceDetail{"Hobbykamer", "Stekker", FibaroWallplug5},
	}
)
