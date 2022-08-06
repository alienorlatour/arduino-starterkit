package main

import (
	"machine"
	"time"
)

const interval = time.Second * 2

func main() {
	tilt := machine.D8
	tilt.Configure(machine.PinConfig{machine.PinInput})

	leds := []machine.Pin{
		machine.D2,
		machine.D3,
		machine.D4,
		machine.D5,
		machine.D6,
		machine.D7,
	}
	for i := range leds {
		leds[i].Configure(machine.PinConfig{Mode: machine.PinOutput})
		leds[i].Set(true)
		time.Sleep(time.Millisecond * 250)
		leds[i].Set(false)
	}

	var previousTime, currentTime time.Time
	previousTime = time.Now()

	nextLitLed := 0
	// prevState := tilt.Get()

	for {
		currentTime = time.Now()

		if currentTime.Sub(previousTime) > interval {
			previousTime = currentTime
			leds[nextLitLed].Set(true)
			nextLitLed++

			if nextLitLed > 6 {
				timeOut(leds)
			}
		}

		// if prevState != tilt.Get() {
		// 	// reset(leds)
		// 	nextLitLed = 0
		// 	prevState = !prevState
		// }
	}

}

func timeOut(leds []machine.Pin) {
	for {
		for i := range leds {
			leds[i].Set(true)
		}
		time.Sleep(time.Millisecond * 250)
		for i := range leds {
			leds[i].Set(false)
		}
		time.Sleep(time.Millisecond * 250)
	}
}
