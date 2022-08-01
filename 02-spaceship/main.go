package main

import (
	"machine"
	"time"
)

// blinkTime represent the time for each state of the blinking light.
const blinkTime = time.Millisecond * 250

func main() {

	// set up the button
	button := machine.D2
	button.Configure(machine.PinConfig{Mode: machine.PinInput})

	// set up the lights
	red1 := machine.D5
	red1.Configure(machine.PinConfig{Mode: machine.PinOutput})
	red2 := machine.D4
	red2.Configure(machine.PinConfig{Mode: machine.PinOutput})
	green := machine.D3
	green.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		if button.Get() {
			// if the button is pressed, turn on the green
			green.Set(true)
			red1.Set(false)
			red2.Set(false)
		} else {
			// blink the red lights
			green.Set(false)
			red1.Set(true)
			red2.Set(true)
			time.Sleep(blinkTime)
			red1.Set(false)
			red2.Set(false)
			time.Sleep(blinkTime)
		}
	}
}
