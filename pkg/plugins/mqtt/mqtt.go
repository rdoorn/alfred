package mqtt

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	"github.com/rdoorn/alfred/internal/integration"
)

type Handler struct {
	pub           *mqtt.Client
	sub           *mqtt.Client
	id            string
	version       string
	debug         bool
	topics        []string
	messageReader func(client mqtt.Client, msg mqtt.Message)
	messageWriter func(topic string, qos byte, retained bool, payload interface{}) error
	onConnect     func(client mqtt.Client)
	eventbus      <-chan integration.Event
}

const (
	NAME    = "mqtt"
	VERSION = "1.0"
)

// New initializes the plugin
func New(opts ...Option) *Handler {
	clientID := uuid.New()
	h := &Handler{
		id:       clientID.String(),
		version:  VERSION,
		eventbus: make(<-chan integration.Event),
	}

	for _, opt := range opts {
		opt(h)
	}

	return h
}

// Name returns the name of the plugin
func (h *Handler) Name() string {
	return NAME
}

// Version returns the version of the plugin
func (h *Handler) Version() string {
	return VERSION
}

// Close closes all connections and channels
func (h *Handler) Shutdown() {
	if h.sub != nil {
		log.Printf("closing sub")
		client := *h.sub
		client.Disconnect(250)
	}
	if h.pub != nil {
		log.Printf("closing pub")
		client := *h.pub
		client.Disconnect(250)
	}
	log.Printf("closed")
}

func (h *Handler) Subscribe(e <-chan integration.Event) error {

	if h.messageReader == nil {
		h.messageReader = h.defaultMessageReader
	}

	if h.messageWriter == nil {
		h.messageWriter = h.defaultMessageWriter
	}

	if h.onConnect == nil {
		h.onConnect = h.defaultOnConnect
	}

	err := h.subscribeWithOnConnect(h.onConnect)
	if err != nil {
		return err
	}

	return nil
}

func (h *Handler) Publish(e *integration.Event) error {
	return h.messageWriter("bla", 0, false, "")
}

func (h *Handler) publish(topic string, qos byte, retained bool, payload interface{}) error {
	if h.pub == nil {
		log.Printf("mqtt: connecting to pub")
		c, err := client(fmt.Sprintf("%s_pub", h.id), nil)
		if err != nil {
			return err
		}
		h.pub = &c
	}

	log.Printf("mqtt: publishing to %s: %v", topic, payload)
	if token := (*h.pub).Publish(topic, qos, retained, payload); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

func (h *Handler) subscribeWithOnConnect(onConnect func(c mqtt.Client)) error {
	log.Printf("mqtt: subscribeWithOnConnect")
	if h.sub == nil {
		log.Printf("mqtt: connecting to sub: %+v", h)
		c, err := client(fmt.Sprintf("%s_sub", h.id), onConnect)
		if err != nil {
			log.Printf("mqtt: error...")
			return err
		}
		log.Printf("connected to sub %+v", c)
		h.sub = &c
	}
	return nil
}

func client(clientID string, onConnect func(c mqtt.Client)) (mqtt.Client, error) {
	log.Printf("mqtt: client")
	urlVar, ok := os.LookupEnv("MQTT_URL")
	if !ok {
		return nil, errors.New("missing environment key: MQTT_URL")
	}

	uri, err := url.Parse(urlVar)
	if err != nil {
		return nil, err
	}

	opts := clientOptions(clientID, uri)
	opts.OnConnect = onConnect
	client := mqtt.NewClient(opts)
	log.Printf("mqtt: client connect")
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	log.Printf("mqtt: client token")
	if err := token.Error(); err != nil {
		log.Printf("mqtt: client err")
		return nil, err
	}
	log.Printf("mqtt: client ok")
	return client, nil
}

/*

type Handler struct {
	pub *mqtt.Client
	sub *mqtt.Client
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Publish(clientID, topic string, qos byte, retained bool, payload interface{}) error {
	if h.pub == nil {
		log.Printf("mqtt: connecting to pub")
		c, err := newclient(fmt.Sprintf("%s_pub", clientID), nil)
		if err != nil {
			return err
		}
		h.pub = &c
	}

	log.Printf("mqtt: publishing to %s: %v", topic, payload)
	if token := (*h.pub).Publish(topic, qos, retained, payload); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

func (h *Handler) Subscribe(clientID, topic string, qos byte, messageHandler func(client mqtt.Client, msg mqtt.Message)) error {

	s := func(c mqtt.Client) {
		log.Printf("mqtt: subscribing to %s", topic)
		if token := c.Subscribe(topic, qos, messageHandler); token.WaitTimeout(3*time.Second) && token.Error() != nil {
			panic(token.Error())
		}
	}

	if h.sub == nil {
		log.Printf("mqtt: connecting to sub")
		c, err := newclient(fmt.Sprintf("%s_sub", clientID), s)
		if err != nil {
			return err
		}
		h.sub = &c
	}
	return nil
}

func (h *Handler) SubscribeWithOnConnect(clientID string, onConnect func(c mqtt.Client)) error {

	if h.sub == nil {
		log.Printf("mqtt: connecting to sub")
		c, err := newclient(fmt.Sprintf("%s_sub", clientID), onConnect)
		if err != nil {
			return err
		}
		h.sub = &c
	}
	return nil
}

func newclient(clientID string, onConnect func(c mqtt.Client)) (mqtt.Client, error) {
	urlStr := getURL()

	uri, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	opts := createClientOptions(clientID, uri)
	opts.OnConnect = onConnect
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		return nil, err
	}
	return client, nil
}

func getURL() string {
	mqttURL, ok := os.LookupEnv("MQTT_URL")
	if !ok {
		panic("missing environment key: MQTT_URL")
	}

	log.Printf("MQTT_URL: %s", mqttURL)
	return mqttURL
}

func Connect(clientId string, urlStr string, onconnect func(c mqtt.Client)) (mqtt.Client, error) {
	uri, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	opts := createClientOptions(clientId, uri)
	opts.OnConnect = onconnect
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		return nil, err
	}
	return client, nil
}
*/

func clientOptions(clientId string, uri *url.URL) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", uri.Host))
	opts.SetUsername(uri.User.Username())
	password, _ := uri.User.Password()
	opts.SetPassword(password)
	opts.SetClientID(clientId)
	return opts
}
