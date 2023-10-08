package main

import (
	"time"

	"machine"
)

const (
	thresholdLow  = 20.0
	thresholdMid  = 23.0
	thresholdHigh = 26.0
)

func main() {
	cpnts := setup()

	cpnts.red1.Set(true)
	time.Sleep(time.Millisecond * 200)
	cpnts.red1.Set(false)
	time.Sleep(time.Millisecond * 200)
	cpnts.red1.Set(true)
	time.Sleep(time.Millisecond * 200)

	for {
		temperature := cpnts.Temperature()

		cpnts.red1.Set(temperature > thresholdLow)
		cpnts.red2.Set(temperature > thresholdMid)
		cpnts.red3.Set(temperature > thresholdHigh)

		time.Sleep(time.Millisecond * 100)
	}
}

func setup() components {
	// set up the lights
	cpnts := components{
		tpt:  machine.ADC{machine.ADC1},
		red1: machine.D4,
		red2: machine.D3,
		red3: machine.D2,
	}

	cpnts.red1.Configure(machine.PinConfig{Mode: machine.PinOutput})
	cpnts.red2.Configure(machine.PinConfig{Mode: machine.PinOutput})
	cpnts.red3.Configure(machine.PinConfig{Mode: machine.PinOutput})

	// set up the temperature
	cpnts.tpt.Configure(machine.ADCConfig{})

	return cpnts
}

type components struct {
	tpt              machine.ADC
	red1, red2, red3 machine.Pin
}

func (c *components) Temperature() float32 {
	voltage := float32(c.tpt.Get()) / 1024 * 5
	temp := (voltage - 0.5) * 100

	println("voltage: %d, temp: %d", voltage, temp)

	return temp
}
