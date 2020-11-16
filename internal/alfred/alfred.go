package alfred

import (
	"log"

	"github.com/rdoorn/alfred/internal/integration"
)

type Handler struct {
	plugins map[string]integration.Plugin // Plugins (can be either reader or writer)
	//readers map[string]integration.Plugin // readers
	//writers map[string]interface{}        // writers
	debug  bool // enable or disable debug logging
	events <-chan integration.Event
}

func New(opts ...Option) *Handler {
	h := &Handler{
		plugins: make(map[string]integration.Plugin),
		//plugins: make(map[string]string),
		//readers: make(map[string]integration.Plugin),
		//writers: make(map[string]interface{}),
		events: make(<-chan integration.Event),
		debug:  false,
	}

	// execute options
	for _, opt := range opts {
		opt(h)
	}

	// subscribe to all there is
	go h.Subscribe()

	return h
}

func (h *Handler) Subscribe() {
	log.Printf("alfred: (h) Subscribe")
	for _, p := range h.plugins {
		log.Printf("subscribing to plugin %s", p)
		if err := p.Subscribe(h.events); err != nil {
			panic(err)
		}

	}
}

func (h *Handler) Shutdown() {
	for _, p := range h.plugins {
		p.Shutdown()
	}
}

func (h *Handler) Events() <-chan integration.Event {
	return h.events
}
