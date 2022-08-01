package main

import (
	"time"

	"machine"
)

const blinkTime = time.Millisecond * 250

func main() {
	button := machine.D2
	button.Configure(machine.PinConfig{Mode: machine.PinInput})

	red1 := machine.D5
	red1.Configure(machine.PinConfig{Mode: machine.PinOutput})
	red2 := machine.D4
	red2.Configure(machine.PinConfig{Mode: machine.PinOutput})
	green := machine.D3
	green.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		if button.Get() {
			green.Set(true)
			red1.Set(false)
			red2.Set(false)
		} else {
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
