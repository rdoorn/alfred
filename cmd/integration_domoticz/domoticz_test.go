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

	// in test always quit after 10 seconds
	timer := time.NewTimer(10 * time.Second)

	// wait for sigint or sigterm for cleanup - note that sigterm cannot be caught
	sigterm := make(chan os.Signal, 10)
	signal.Notify(sigterm, os.Interrupt, syscall.SIGTERM)

	for {
		select {
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
