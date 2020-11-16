package main

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	pahomqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/rdoorn/alfred/internal/alfred"
	"github.com/rdoorn/alfred/pkg/plugins/mqtt"
)

type domoticz struct{}

func main() {
	log.Printf("Starting")
	d := &domoticz{}

	handler := alfred.New(
		alfred.Subscribe(
			mqtt.New(
				mqtt.Debug(true),
				mqtt.Topic("domoticz/+"),
				mqtt.MessageReader(d.ConvertToAlfred),
			),
		),
	)

	// wait for sigint or sigterm for cleanup - note that sigterm cannot be caught
	sigterm := make(chan os.Signal, 10)
	signal.Notify(sigterm, os.Interrupt, syscall.SIGTERM)

	log.Printf("integration domoticz: handler start")
	for {
		select {
		case <-sigterm:
			log.Fatal("sigterm received")
			handler.Shutdown()
			os.Exit(1)
		}
	}
}

func (d *domoticz) ConvertToAlfred(client pahomqtt.Client, msg pahomqtt.Message) {
	log.Printf("mqtt: * [%s] %s\n", msg.Topic(), strings.ReplaceAll(string(msg.Payload()), "\n", ""))

	event, err := parse(msg.Payload())
	if err != nil {
		log.Printf("domoticz/out failed to process payload: %s", err)
	}

	log.Printf("event: %+v", event)
	/*switch msg.Topic() {
	case "domoticz/out":
		d, err := mqttdomoticz.Parse(msg.Payload())
		if err != nil {
			log.Printf("domoticz/out failed to process payload: %s", err)
		}
	}*/

}

type state struct {
	Battery        int    `json:"Battery"`
	RSSI           int    `json:"RSSI"`
	Description    string `json:"description"`
	DeviceType     string `json:"dtype"`
	DeviceID       string `json:"deviceid"`
	DeviceLocation string `json:"devicelocation"`
	DeviceName     string `json:"devicename"`
	DeviceHardware int    `json:"devicehardware"`
	ID             string `json:"id"`
	Index          int    `json:"idx"`
	Name           string `json:"name"`
	NormalValue    int    `json:"nvalue"`
	StringType     string `json:"stype"`
	StringValue1   string `json:"svalue1"`
	StringValue2   string `json:"svalue2"`
	Unit           int    `json:"unit"`
}

func parse(payload []byte) (*state, error) {
	v := &state{}
	if err := json.Unmarshal(payload, v); err != nil {
		return nil, err
	}
	if v.ID[len(v.ID)-4:][:2] == "00" {
		v.DeviceID = v.ID[len(v.ID)-2:][:2]
	} else {
		v.DeviceID = v.ID[len(v.ID)-4:][:2]
	}
	if _, ok := ZwaveDeviceDetails[v.DeviceID]; ok {
		v.DeviceLocation = ZwaveDeviceDetails[v.DeviceID].Location
		v.DeviceName = ZwaveDeviceDetails[v.DeviceID].Name
		v.DeviceHardware = ZwaveDeviceDetails[v.DeviceID].Type
	}
	return v, nil
}
