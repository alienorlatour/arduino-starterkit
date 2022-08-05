package main

import (
	"machine"
	"time"
)

func main() {
	vehic := newSemaphore(machine.D12, machine.D11, machine.D10)
	pedestr := newPedestrian(machine.D7, machine.D6, machine.D5)

	button := machine.D2
	button.Configure(machine.PinConfig{Mode: machine.PinInput})

	vehic.green.Set(true)

	for {
		if button.Get() {
			// somebody pressed the button
			pedestr.SignalKommt()

			vehic.RedGerman()
			pedestr.Green()

			// pedestrians cross
			time.Sleep(time.Second * 4)

			pedestr.Red()
			vehic.Green()
		}

		time.Sleep(time.Millisecond * 25)
	}

}

type Pedestrian struct {
	red, green, signal machine.Pin
}

func newPedestrian(r, g, signal machine.Pin) Pedestrian {
	r.Configure(machine.PinConfig{Mode: machine.PinOutput})
	g.Configure(machine.PinConfig{Mode: machine.PinOutput})
	signal.Configure(machine.PinConfig{Mode: machine.PinOutput})

	return Pedestrian{r, g, signal}
}

func (s Pedestrian) SignalKommt() {
	for i := 0; i < 10; i++ {
		s.signal.Set(true)
		time.Sleep(time.Millisecond * 250)

		s.signal.Set(false)
		time.Sleep(time.Millisecond * 250)
	}
}

func (s Pedestrian) Green() {
	s.red.Set(false)
	s.signal.Set(false)
	s.green.Set(true)
}

func (s Pedestrian) Red() {
	s.signal.Set(false)

	// blink the green
	for i := 0; i < 5; i++ {
		s.green.Set(true)
		time.Sleep(time.Millisecond * 250)

		s.green.Set(false)
		time.Sleep(time.Millisecond * 250)
	}

	s.red.Set(true)
}

func newSemaphore(r, y, g machine.Pin) Semaphore {
	r.Configure(machine.PinConfig{Mode: machine.PinOutput})
	y.Configure(machine.PinConfig{Mode: machine.PinOutput})
	g.Configure(machine.PinConfig{Mode: machine.PinOutput})

	return Semaphore{
		red:    r,
		yellow: y,
		green:  g,
	}
}

type Semaphore struct {
	red    machine.Pin
	yellow machine.Pin
	green  machine.Pin
}

func (t *Semaphore) Green() {
	t.red.Set(true)
	t.yellow.Set(true)
	t.green.Set(false)

	time.Sleep(time.Second)

	t.yellow.Set(false)
	t.red.Set(false)
	t.green.Set(true)
}

func (t *Semaphore) RedGerman() {
	t.red.Set(false)
	t.yellow.Set(true)
	t.green.Set(false)

	time.Sleep(time.Second)

	t.yellow.Set(false)
	t.red.Set(true)
}
