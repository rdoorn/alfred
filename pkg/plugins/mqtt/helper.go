package mqtt

import (
	"log"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func (h *Handler) defaultOnConnect(c mqtt.Client) {
	log.Printf("mqtt: using defaultOnConnect")
	// subscribe to channels
	for _, topic := range h.topics {
		log.Printf("mqtt: subscribing to %s", topic)
		if token := c.Subscribe(topic, 0, h.messageReader); token.WaitTimeout(3*time.Second) && token.Error() != nil {
			panic(token.Error())
		}
	}
	// execute inits
	//for _, init := range h.SubscriptionInits {
	//c.Publish(init[0], 0, false, init[1])
	//}

}

func (h *Handler) defaultMessageReader(client mqtt.Client, msg mqtt.Message) {
	log.Printf("mqtt: * [%s] %s\n", msg.Topic(), strings.ReplaceAll(string(msg.Payload()), "\n", ""))

	/*switch msg.Topic() {
	case "domoticz/out":
		d, err := mqttdomoticz.Parse(msg.Payload())
		if err != nil {
			log.Printf("domoticz/out failed to process payload: %s", err)
		}
	}*/

}

func (h *Handler) defaultMessageWriter(topic string, qos byte, retained bool, payload interface{}) error {
	return h.publish(topic, qos, retained, payload)
}
