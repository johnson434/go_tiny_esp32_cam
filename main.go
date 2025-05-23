package main

import (
	"fmt"
	"machine"
	"time"
)

func main() {
	fmt.Print("Hello World")
	led := machine.GPIO4
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		led.High()
		time.Sleep(time.Second / 2)

		led.Low()
		time.Sleep(time.Second / 2)
	}
}
