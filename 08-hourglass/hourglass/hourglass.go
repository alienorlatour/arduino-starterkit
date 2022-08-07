package hourglass

import (
	"machine"
	"time"
)

const interval = time.Second * 1

func NewHourglass(
	pins []machine.Pin,
	turn chan bool,
) *Hourglass {
	for i := range pins {
		pins[i].Configure(machine.PinConfig{Mode: machine.PinOutput})
	}

	return &Hourglass{
		leds: pins,
		turn: turn,
	}
}

type Hourglass struct {
	leds      []machine.Pin
	nextLight int
	turn      chan bool
}

func (h *Hourglass) Run() {
	previousTime, currentTime := time.Now(), time.Now()

	for {
		currentTime = time.Now()

		if currentTime.Sub(previousTime) > interval {
			previousTime = currentTime
			timesUp := h.next()

			if timesUp {
				// todo
				// time
				// 's up
			}
		}

		select {
		case _ = <-h.turn:
			previousTime = currentTime
			h.reset()
		default:
		}

		time.Sleep(time.Millisecond * 25)
	}
}

func (h *Hourglass) reset() {
	for i := range h.leds {
		h.leds[i].Set(false)
	}
	h.nextLight = 0
}

func (h *Hourglass) next() bool {
	if h.nextLight >= len(h.leds) {
		return true
	}
	h.leds[h.nextLight].Set(true)
	h.nextLight++
	return false
}
