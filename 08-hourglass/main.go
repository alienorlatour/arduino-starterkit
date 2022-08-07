package main

import (
	"machine"
	"time"

	"github.com/ablqk/arduino-starterkit/08-hourglass/hourglass"
)

func main() {
	tilt := machine.D8
	tilt.Configure(machine.PinConfig{machine.PinInput})

	// control the state of the tilt switch with the built-in LED
	tiltState := machine.D13
	tiltState.Configure(machine.PinConfig{Mode: machine.PinOutput})

	turn := make(chan bool)

	// create an hourglass
	h := hourglass.NewHourglass(
		[]machine.Pin{
			machine.D2,
			machine.D3,
			machine.D4,
			machine.D5,
			machine.D6,
			machine.D7,
		},
		turn,
	)

	go h.Run()

	prevState := tilt.Get()

	for {
		tiltState.Set(tilt.Get())

		if prevState != tilt.Get() {
			prevState = !prevState
			turn <- prevState
		}

		time.Sleep(time.Millisecond * 25)
	}

}

func timeOut() {

}
