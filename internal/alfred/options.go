package alfred

import (
	"log"

	"github.com/rdoorn/alfred/internal/integration"
)

type Option func(*Handler)

func Debug() Option {
	return func(h *Handler) {
		h.debug = true
	}
}

func Subscribe(i integration.Plugin) Option {
	return func(h *Handler) {
		name := i.Name()
		log.Printf("loading plugin %s", name)
		if val, ok := h.plugins[name]; ok {
			log.Printf("plugin: %s already loaded", val)
		}
		h.plugins[name] = i
	}
}
