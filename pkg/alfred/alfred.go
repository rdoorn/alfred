package alfred

import (
	"fmt"
	"log"

	"github.com/rdoorn/alfred/pkg/zwave"
)

const (
	version = "0.1.2"
)

type Handler struct {
	config   Config
	zwave    []*zwave.Controller
	http     *HttpHandler
	shutdown chan struct{}
}

func New() (*Handler, error) {
	h := &Handler{
		http:     NewHttpHandler(),
		shutdown: make(chan struct{}),
	}

	// load configuration
	if err := h.LoadConfig(); err != nil {
		return nil, fmt.Errorf("failed to load config file %s, error: %s", configFile, err)
	}

	// setup http interfaces
	v1 := h.http.handler.Group("/v1")
	{
		v1.GET("/version", h.version)
		v1.GET("/nodes", h.nodes)
		v1.GET("/nodes/:id", h.node)
		//v1.GET("/hello/:name", handler.hello)
	}
	go h.http.handler.Run(fmt.Sprintf("%s:%d", h.config.Listeners.Http.Listener, h.config.Listeners.Http.Port))

	// start devices
	for id, dev := range h.config.Devices.Zwave {
		log.Printf("starting zwave listener %s at %d", dev.Device, dev.Baudrate)
		zw, err := zwave.New(zwave.WithDevice(zwave.NewSerial(dev.Device, dev.Baudrate)))
		zw.DiscoverNodes()
		if err != nil {
			return nil, fmt.Errorf("failed to start zwave device %d, error: %s", id, err)
		}
		h.zwave = append(h.zwave, zw)
	}
	// start routing messages between devices
	go h.Router()

	log.Printf("handler: %+v", h)
	return h, nil
}

func (h *Handler) Shutdown() {
	close(h.shutdown)
}

func (h *Handler) Router() {
	log.Printf("started router")
	for {
		select {
		case _ = <-h.shutdown:
			log.Printf("stopping router")
			return
		}
	}
}
