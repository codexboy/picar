package electrical

import (
	"github.com/stianeikeland/go-rpio"
)

const (
	SStop = iota
	SForward
	SBackup
)

type Elect struct {
	pin1  rpio.Pin
	pin2  rpio.Pin
	State uint8
}

func NewElect(pin1, pin2 uint8) *Elect {
	e := Elect{pin1: rpio.Pin(pin1), pin2: rpio.Pin(pin2), State: SStop}
	e.pin1.Output()
	e.pin2.Output()
	e.Stop()
	return &e
}

func (e *Elect) Forward() {
	e.pin1.High()
	e.pin2.Low()
	e.State = SForward
}

func (e *Elect) Stop() {
	e.pin1.Low()
	e.pin2.Low()
	e.State = SStop
}

func (e *Elect) Backup() {
	e.pin1.Low()
	e.pin2.High()
	e.State = SBackup
}
