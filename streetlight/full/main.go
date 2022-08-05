package main

import (
	"machine"
	"time"
)

func main() {
	tri := newTricolor(machine.D12, machine.D11, machine.D10)
	pedestr := newPedestrian(machine.D7, machine.D6, machine.D5)

	button := machine.D2
	button.Configure(machine.PinConfig{Mode: machine.PinInput})

	tri.green.Set(true)

	for {
		if button.Get() {
			// somebody pressed the button
			pedestr.SignalKommt(10)

			tri.TransRedGerman()
			pedestr.Green()

			// pedestrians cross
			time.Sleep(time.Second * 6)

			pedestr.TranRedFr()
			tri.TransGreen()
		}

		time.Sleep(time.Millisecond * 25)
	}

}

type Pedestrian struct {
	r, g, signal machine.Pin
}

func newPedestrian(r, g, signal machine.Pin) Pedestrian {
	r.Configure(machine.PinConfig{Mode: machine.PinOutput})
	g.Configure(machine.PinConfig{Mode: machine.PinOutput})
	signal.Configure(machine.PinConfig{Mode: machine.PinOutput})

	return Pedestrian{r, g, signal}
}

func (s Pedestrian) SignalKommt(times int) {
	for i := 0; i < times; i++ {
		s.signal.Set(true)
		time.Sleep(time.Millisecond * 250)

		s.signal.Set(false)
		time.Sleep(time.Millisecond * 250)
	}
}

func (s Pedestrian) Green() {
	s.r.Set(false)
	s.signal.Set(false)
	s.g.Set(true)
}

func (s Pedestrian) TranRedFr() {
	s.signal.Set(false)

	// blink the green
	for i := 0; i < 5; i++ {
		s.g.Set(true)
		time.Sleep(time.Millisecond * 250)

		s.g.Set(false)
		time.Sleep(time.Millisecond * 250)
	}

	s.r.Set(true)
}

func newTricolor(r, y, g machine.Pin) Tricolor {
	t := Tricolor{
		red:    r,
		yellow: y,
		green:  g,
	}

	t.red.Configure(machine.PinConfig{Mode: machine.PinOutput})
	t.yellow.Configure(machine.PinConfig{Mode: machine.PinOutput})
	t.green.Configure(machine.PinConfig{Mode: machine.PinOutput})

	return t
}

type Tricolor struct {
	red    machine.Pin
	yellow machine.Pin
	green  machine.Pin
}

func (t *Tricolor) TransGreen() {
	t.red.Set(true)
	t.yellow.Set(true)
	t.green.Set(false)

	time.Sleep(time.Second)

	t.yellow.Set(false)
	t.red.Set(false)
	t.green.Set(true)
}

func (t *Tricolor) TransRedGerman() {
	t.red.Set(false)
	t.yellow.Set(true)
	t.green.Set(false)

	time.Sleep(time.Second)

	t.yellow.Set(false)
	t.red.Set(true)
}

//
// func (t *Tricolor) Run(passTime time.Duration) {
// 	for {
// 		t.yellow.Set(false)
// 		t.red.Set(false)
// 		t.green.Set(true)
//
// 		time.Sleep(passTime)
//
// 		t.green.Set(false)
// 		t.yellow.Set(true)
//
// 		time.Sleep(time.Second)
//
// 		t.yellow.Set(false)
// 		t.red.Set(true)
//
// 		time.Sleep(passTime)
//
// 		t.yellow.Set(true)
//
// 		time.Sleep(time.Second)
// 	}
// }
