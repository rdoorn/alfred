package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"

	"github.com/rdoorn/alfred/internal/alfred"
	"github.com/rdoorn/alfred/pkg/plugins/mqtt"
)

func TestLivingroom(t *testing.T) {
	handler := alfred.New(
		alfred.Subscribe(
			mqtt.New(
				mqtt.Debug(true),
				mqtt.Topic("alfred/livingroom/sensor/+/lux"),
				mqtt.Topic("alfred/livingroom/sensor/+/motion"),
				mqtt.Topic("alfred/livingroom/switch/+/state"),
				mqtt.Topic("alfred/generic/alarm/state"),
				mqtt.Topic("domoticz/+"),
			),
		),
	)

	// in test always quit after 10 seconds
	timer := time.NewTimer(10 * time.Second)

	// wait for sigint or sigterm for cleanup - note that sigterm cannot be caught
	sigterm := make(chan os.Signal, 10)
	signal.Notify(sigterm, os.Interrupt, syscall.SIGTERM)

	for {
		select {
		case event := <-handler.Events():
			log.Printf("event: %+v", event)
		case <-timer.C:
			log.Fatal("timer expired")
			os.Exit(1)
		case <-sigterm:
			handler.Shutdown()
			log.Fatal("sigterm received")
			os.Exit(1)
		}
	}

}
