package mqtt

import mqtt "github.com/eclipse/paho.mqtt.golang"

type Option func(*Handler)

// Debug enable or disabled debug
func Debug(b bool) Option {
	return func(h *Handler) {
		h.debug = b
	}
}

// ClientID overrides the name as which to connect to mqtt (Default: uuid)
func ClientID(s string) Option {
	return func(h *Handler) {
		h.id = s
	}
}

// Topic adds a subscription to a topic (can be used multiple times)
func Topic(s string) Option {
	return func(h *Handler) {
		h.topics = append(h.topics, s)
	}
}

// MessageReader overrides the default message handler for reads
func MessageReader(messageReader func(client mqtt.Client, msg mqtt.Message)) Option {
	return func(h *Handler) {
		h.messageReader = messageReader
	}
}

// MessageWriter overrides the default message handler for reads
func MessageWriter(messageWriter func(topic string, qos byte, retained bool, payload interface{}) error) Option {
	return func(h *Handler) {
		h.messageWriter = messageWriter
	}
}
