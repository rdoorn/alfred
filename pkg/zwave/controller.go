package zwave

import "log"

type Controller struct {
	debug     bool
	device    Device    // the device supporting io
	transport Transport // reads the device, and cuts it in to requests
	//translator Translator // translates the requests to messages
	session Session // work at session level, handle retries and timeouts
}

func New(opts ...Option) (*Controller, error) {
	c := &Controller{
		device:    NewSerial("/dev/tty.usbmodem14101", 115200),
		transport: NewZwaveAPI(),
		//translator: NewTranslator(),
		//session:   NewZwaveSession(nil),
	}

	// defaults
	//WithDevice(NewSerial("/dev/tty.usbmodem14101", 115200))(c)
	//WithTransport(NewZwaveAPI())(c)

	for _, opt := range opts {
		opt(c)
	}

	rw, err := c.device.Open()
	if err != nil {
		log.Fatalf("failed to open device %s error: %s", c.device, err)
	}
	log.Printf("Connected to %s", c.device)

	err = c.transport.Hook(rw)
	if err != nil {
		log.Fatal("Failed to hook transport: err", err)
	}
	log.Printf("Transporter configured: %s", c.transport)

	c.session = NewZwaveSession(c.transport)
	/*if err != nil {
		log.Fatal("Failed to create session: err", err)
	}*/
	log.Printf("Session configured: %s", c.session)

	go c.transport.Reader()
	go c.transport.Writer()
	go c.session.Reader()

	return c, nil

}

func (c *Controller) Shutdown() {

}

func (c *Controller) DiscoverNodes() {
	//c.session.WriteFunction(DiscoveryNodes)
	c.session.Write(SOF(REQ, DiscoveryNodes))
	//	c.session.Write(NewMessage(Request, DiscoveryNodes))

}
