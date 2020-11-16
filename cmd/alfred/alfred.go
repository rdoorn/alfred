package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rdoorn/alfred/pkg/zwave"
)

var port string

func main() {

	zw, err := zwave.New(
	//zwave.WithDevice(zwave.NewSerial("/dev/tty.usbmodem14101", 115200)),
	)
	if err != nil {
		log.Fatal(err)
	}

	zw.DiscoverNodes()

	// in test always quit after 10 seconds
	timer := time.NewTimer(10 * time.Second)

	// wait for sigint or sigterm for cleanup - note that sigterm cannot be caught
	sigterm := make(chan os.Signal, 10)
	signal.Notify(sigterm, os.Interrupt, syscall.SIGTERM)

	log.Printf("zwave handler start")
	for {
		select {
		case <-timer.C:
			log.Fatal("timer expired")
			os.Exit(1)
		case <-sigterm:
			log.Fatal("sigterm received")
			zw.Shutdown()
			os.Exit(1)
		}
	}

	/*
		c := &serial.Config{Name: "/dev/tty.usbmodem14101", Baud: 115200}
		s, err := serial.OpenPort(c)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%T", s)

		time.Sleep(1 * time.Second)
		msg := zwave.NewRequest(0x02)
		log.Printf("writing: %q", msg.Message())
		os.Exit(1)
		b := serialapi.NewRaw(msg.Message())
		log.Printf("writing: %q", b)
		n, err := s.Write(b)
		if err != nil {
			log.Fatal(err)
		}

		buf := make([]byte, 128)
		n, err = s.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("reading: %q", buf[:n])
	*/
}

/*

	r := gin.Default()
	r.GET("/nodes", getNodes(z))
	r.GET("/nodes/:id", getNode(z))
	r.GET("/nodes/:id/:action", control(z))
	go r.Run() // listen and serve on 0.0.0.0:8080

	for {
		select {
		case event := <-z.GetNextEvent():
			log.Println("----------------------------------------")
			log.Printf("Event: %#v\n", event)
			switch e := event.(type) {
			case events.NodeDiscoverd:
				znode := z.Nodes.Get(e.Address)
				znode.RLock()
				log.Printf("Node: %#v\n", znode)
				znode.RUnlock()

			case events.NodeUpdated:
				znode := z.Nodes.Get(e.Address)
				znode.RLock()
				log.Printf("Node: %#v\n", znode)
				znode.RUnlock()
			}
			log.Println("----------------------------------------")
		}
	}
*/
