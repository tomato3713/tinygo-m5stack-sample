package main

import (
	"machine"
	"time"
	"fmt"
	"image/color"
	"github.com/tomato3713/m5tiny"
	"tinygo.org/x/tinyfont/freemono"
)

var blue  = color.RGBA{0, 0, 255, 255}
var black  = color.RGBA{0, 0, 0, 255}


func main() {
	display := m5tiny.InitDisplay()

	btnA := machine.BUTTON_A
	btnB := machine.BUTTON_B
	btnC := machine.BUTTON_C

	btnA.Configure(machine.PinConfig{Mode: machine.PinInput})
	btnB.Configure(machine.PinConfig{Mode: machine.PinInput})
	btnC.Configure(machine.PinConfig{Mode: machine.PinInput})

	w, h := display.Size()
	display.FillRectangle(0, 0, w, h, black)
	ch := ""

	for {
		if !btnA.Get() {
			fmt.Println("pressed A")
			ch = "A"
		} else if !btnB.Get() {
			fmt.Println("pressed B")
			ch = "B"
		} else if !btnC.Get() {
			fmt.Println("pressed C")
			ch = "C"
		} else {
			ch = ""
		}

		display.FillRectangle(0, 20, w, 25, black)
		display.WriteLine(&freemono.Regular9pt7b, 30, 40, "pressed key " + ch, blue)
		time.Sleep(time.Millisecond * 50)	
	}
}