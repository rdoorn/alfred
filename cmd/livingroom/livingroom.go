package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/rdoorn/alfred/internal/alfred"
	"github.com/rdoorn/alfred/internal/integration"
	"github.com/rdoorn/alfred/pkg/plugins/mqtt"
)

const (

)

func main() {
	log.Printf("livingroom: new")
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
	log.Printf("livingroom: handler set")

	// wait for sigint or sigterm for cleanup - note that sigterm cannot be caught
	sigterm := make(chan os.Signal, 10)
	signal.Notify(sigterm, os.Interrupt, syscall.SIGTERM)

	log.Printf("livingroom: handler start")
	for {
		select {
		case event := <-handler.Events():
			log.Printf("event: %+v", event)

			switch event.Type {
			case integration.EventSensorMotion:
				if event.Motion() {
						if handler.Get(1).Lux()

				}

			case integration.EventSensorLux:
				if event.Id() == 1 && event.Lux() < 20 {
					// switch on light
				}
			}

		case <-sigterm:
			log.Fatal("sigterm received")
			handler.Shutdown()
			os.Exit(1)
		}
	}
}

// woonkamer motion 311
// woonkamer lux 313
// woonkamer scene on 1
// woonkamer scene off 17

// woonkamer2 motion 321 (kast)
// woonkamer2 lux 323

// woonkamer lampen: 15, 16, 18 (groen hal, tuin)
// eetmakerspot: 211

/*
var (
	domoticzDevices       = []int{15, 16, 18, 211, 311, 313, 321, 323}
	automaticTimeout      = 20 * time.Minute
	automaticTimeoutShort = 5 * time.Minute
	manualTimeout         = 30 * time.Minute
	buitenCam             = "reolink-buiten1"
)

func main() {
	log.Printf("Livingroom")
	h := eventhandler.New(
		eventhandler.ClientID("mqtt_event_handler_livingroom2"),
		eventhandler.SubscribeDomoticzDevices(domoticzDevices...),
		eventhandler.SubscribeVerisureAlarmState(),
		eventhandler.SubscribeReolinkMotionState(buitenCam),
		//eventhandler.Debug(),
	)

	poller, err := h.Poll()
	if err != nil {
		panic(err)
	}

	for {
		select {
		case event := <-poller:
			if event.Type == eventhandler.EventExit {
				log.Printf("client exiting")
				os.Exit(0)
			}

			if event.Type != eventhandler.EventTimerExpired {
				log.Printf("event: %v", event)
			}

			log.Printf("state: alarm:%s sensor1:%t sensor2:%t lux1:%d lux2:%d timer:%t expired:%t l1:%t l2:%t l3:%t l4:%t reolink-buiten1:%t",
				h.Vrsr.AlarmState(),
				h.Dmtcz.On(311), h.Dmtcz.On(321), h.Dmtcz.Lux(313), h.Dmtcz.Lux(323),
				h.Timer(1), h.TimerExpired(1),
				h.Dmtcz.On(15), h.Dmtcz.On(16), h.Dmtcz.On(18), h.Dmtcz.On(211),
				h.Reolink.Motion(buitenCam),
			)
			// new motion, set scene and timer only if lux is low
			// if timer is not set (e.g. all should be off), or if light 18 is off
			if (h.Dmtcz.On(321) || h.Dmtcz.On(311)) && h.Vrsr.Disarmed() && (!h.Timer(1) || !h.Dmtcz.On(18)) && h.Dmtcz.Lux(313) < 20 {
				h.SetTimer(1, automaticTimeout)
				h.Publish("domoticz/in", 0, false, mqttdomoticz.EnableScene(1))
				continue
			}

			if h.Dmtcz.On(321) && h.Vrsr.ArmedHome() && !h.Timer(1) && h.Dmtcz.Lux(323) < 22 {
				log.Printf("state2")
				h.SetTimer(1, automaticTimeoutShort)
				h.Publish("domoticz/in", 0, false, mqttdomoticz.EnableScene(18))
				continue
			}

			// there was motion and a timer, extend timer
			if (h.Dmtcz.On(321) || h.Dmtcz.On(311)) && h.Timer(1) && h.Dmtcz.Lux(313) < 80 {
				log.Printf("state3")
				h.ResetTimer(1)
				continue
			}

			// the ligts are on, but no timer
			//if (h.Dmtcz.On(15) || h.Dmtcz.On(16) || h.Dmtcz.On(18) || h.Dmtcz.On(211)) && !h.Timer(1) {
			if (h.Dmtcz.On(15) || h.Dmtcz.On(18) || h.Dmtcz.On(211)) && !h.Timer(1) {
				log.Printf("state6")
				h.SetTimer(1, manualTimeout)
			}

			// the ligts are off, but still a timer running
			if !(h.Dmtcz.On(15) || h.Dmtcz.On(16) || h.Dmtcz.On(18) || h.Dmtcz.On(211)) && h.Timer(1) {
				log.Printf("state6b")
				h.DeleteTimer(1)
			}

			// When we are not home, and someone gets near in the dark, light a light
			if h.Reolink.Motion(buitenCam) && h.Vrsr.ArmedAway() && h.Dmtcz.Lux(313) < 20 {
				log.Printf("stateAway1")
				h.SetTimer(1, automaticTimeoutShort)
				h.Publish("domoticz/in", 0, false, mqttdomoticz.SetBrightness(16, 100))
				continue
			}

			// timer expired and no motion, disable scene
			if h.Timer(1) && h.TimerExpired(1) {
				log.Printf("state7")
				h.DeleteTimer(1)
				h.Publish("domoticz/in", 0, false, mqttdomoticz.EnableScene(17))
			}

			if h.Timer(1) {
				log.Printf("state8")
				log.Printf("timer1 expiry in: %.1fs", h.TimerDuration(1).Minutes())
			}

		}
	}

}

*/
