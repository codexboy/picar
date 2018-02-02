package infrared

import (
	"github.com/stianeikeland/go-rpio"
)

type Infrared struct {
	Pin rpio.Pin
}

func NewInfrared(pin uint8) *Infrared {
	ipin := Infrared{Pin:rpio.Pin(pin)}
	ipin.Pin.Input()

	return &ipin;
}

func (i *Infrared) Check() bool {
	stat := i.Pin.Read();
	if stat == rpio.Low {
		return true;
	} else {
		return false;
	}
}
