package main

//
// func main() {
//
//	machine.InitADC()                // initialize ADC
//	ldr := machine.ADC{machine.ADC0} // set ADC pin
//	// same as machine.ADC{machine.ADC0}
//	ldr.Configure(machine.ADCConfig{}) // start the pin's ADC function
//
//	led := machine.LED
//	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
//
//	//var blinker bool
//
//	for {
//		// output analog reading to serial port
//		print(ldr.Get())
//		// light up led if analog reading is over the threshold
//		led.Set(ldr.Get() > 40000)
//		delay(100)
//	}
// }
//
// func delay(t int64) {
//	time.Sleep(time.Duration(1000000 * t))
// }
