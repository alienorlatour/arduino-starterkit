package main

import (
	"machine"
	"time"
)

type rgbLED struct {
	rCh, gCh, bCh uint8
	pwm           machine.PWM
}

func newRGB(r, g, b machine.Pin) (*rgbLED, error) {
	pwm := machine.Timer1

	err := pwm.Configure(machine.PWMConfig{})
	if err != nil {
		return nil, err
	}

	chR, err := pwm.Channel(r)
	if err != nil {
		return nil, err
	}

	chG, err := pwm.Channel(g)
	if err != nil {
		return nil, err
	}

	chB, err := pwm.Channel(b)
	if err != nil {
		return nil, err
	}

	return &rgbLED{
		rCh: chR,
		gCh: chG,
		bCh: chB,
		pwm: pwm,
	}, nil
}

func (r *rgbLED) SetSpectrum(red, green, blue uint32) {
	r.pwm.Set(r.rCh, red)
	r.pwm.Set(r.gCh, green)
	r.pwm.Set(r.bCh, blue)
}

type sensor struct {
	r, g, b machine.ADC
}

func newSensor(r, g, b machine.Pin) *sensor {
	config := machine.ADCConfig{}

	red := machine.ADC{Pin: r}
	red.Configure(config)

	green := machine.ADC{Pin: g}
	green.Configure(config)

	blue := machine.ADC{Pin: b}
	blue.Configure(config)

	return &sensor{
		r: red,
		g: green,
		b: blue,
	}
}

func (s *sensor) ReadSpectrum() (r, g, b uint32) {
	return uint32(1024 - s.r.Get()),
		uint32(1024 - s.g.Get()),
		uint32(1024 - s.b.Get())
}

func main() {

	rgb, err := newRGB(machine.D11, machine.D9, machine.D10)
	if err != nil {
		println(err)
		return
	}

	sensor := newSensor(machine.ADC0, machine.ADC1, machine.ADC3)

	for {

		rgb.SetSpectrum(sensor.ReadSpectrum())

		time.Sleep(time.Millisecond * 25)
	}
}
