package alfred

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var (
	configFile string = "config.yaml"
)

type Config struct {
	Devices   DevicesConfig   `yaml:"devices"`
	Listeners ListenersConfig `yaml:"listeners"`
}

type DevicesConfig struct {
	Zwave []ZwaveConfig `yaml:"zwave"`
}

type ZwaveConfig struct {
	Device   string `yaml:"device"`
	Baudrate int    `yaml:"baudrate"`
}

type ListenersConfig struct {
	Http HttpConfig `yaml:"http"`
}

type HttpConfig struct {
	Listener string
	Port     int
	Port2    int
	// TLS cert
}

func (h *Handler) LoadConfig() error {

	/*
		t := Config{
			Devices: DevicesConfig{
				Zwave: []ZwaveConfig{
					ZwaveConfig{
						Device:   "device",
						Baudrate: 1234,
					},
				},
			},
		}
		//f, _ := yaml.Marshal(t)
		//log.Printf("%s", f)
		//ioutil.WriteFile(configFile, f, 0644)
	*/

	yamlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}

	// Set default
	config := &Config{
		Listeners: ListenersConfig{
			Http: HttpConfig{
				Listener: "127.0.0.1",
				Port:     80,
			},
		},
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		return err
	}
	h.config = *config
	return nil
}
