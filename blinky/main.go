package main

import (
	"time"

	"machine"
)

const blinkTime = time.Millisecond * 500

func main() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		led.Set(false)
		time.Sleep(blinkTime)

		led.Set(true)
		time.Sleep(blinkTime)
	}
}
