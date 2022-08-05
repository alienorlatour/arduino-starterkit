package main

import (
	"machine"
	"time"
)

func main() {
	tri := newSemaphore(machine.D12, machine.D11, machine.D10)
	tri.Run(time.Second * 3)

	// signalKommt := machine.D5
	// signalKommt.Configure(machine.PinConfig{Mode: machine.PinOutput})
	// blink(signalKommt, time.Millisecond*500)
}

// func blink(pin machine.Pin, duration time.Duration) {
// 	for {
// 		pin.Set(true)
// 		time.Sleep(duration)
//
// 		pin.Set(false)
// 		time.Sleep(duration)
// 	}
// }

type semaphore struct {
	red    machine.Pin
	yellow machine.Pin
	green  machine.Pin
}

func newSemaphore(r, y, g machine.Pin) semaphore {
	t := semaphore{
		red:    r,
		yellow: y,
		green:  g,
	}

	t.red.Configure(machine.PinConfig{Mode: machine.PinOutput})
	t.yellow.Configure(machine.PinConfig{Mode: machine.PinOutput})
	t.green.Configure(machine.PinConfig{Mode: machine.PinOutput})

	return t
}

func (t *semaphore) Run(passTime time.Duration) {
	for {
		// t.yellow.Set(false)
		t.red.Set(false)
		t.green.Set(true)

		time.Sleep(passTime)

		t.green.Set(false)
		// t.yellow.Set(true)

		// time.Sleep(time.Second)

		// t.yellow.Set(false)
		t.red.Set(true)

		time.Sleep(passTime)

		// t.yellow.Set(true)
		// time.Sleep(time.Second)
	}
}
