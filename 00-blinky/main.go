package main

import (
	"machine"
	"time"
)

const blinkTime = time.Millisecond * 1000

func main() {
	// set built-in LED as output
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	// blink
	for {
		led.Set(false)
		time.Sleep(blinkTime)

		led.Set(true)
		time.Sleep(blinkTime)
	}
}
