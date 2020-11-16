package zwave

type Option func(*Controller) error

func Debug() Option {
	return func(h *Controller) error {
		h.debug = true
		return nil
	}
}

func WithDevice(d Device) Option {
	return func(h *Controller) error {

		h.device = d

		/*
			rw, err := d.Open()
			if err != nil {
				return fmt.Errorf("Device failure: %s", err)
			}
			log.Printf("Connected to %s", d)
			h.rw = rw
			return nil
		*/
		return nil
	}
}

func WithTransport(t Transport) Option {
	return func(h *Controller) error {

		/*
				t, err := d.Hook(h.device.Open())
				if err != nil {
					return fmt.Errorf("Transport failure: %s", err)
				}
				log.Printf("Transport setup to use %T", d)
				h.transport = t
				return nil
			}
		*/

		h.transport = t
		return nil
	}
}
