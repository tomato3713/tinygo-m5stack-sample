package main

import (
	"github.com/tomato3713/m5tiny"
	"tinygo.org/x/tinyfont/freemono"
	"image/color"
)

var (
	display *m5tiny.DisplayDevice
)

var (
	black = color.RGBA{0, 0, 0, 255}
	blue  = color.RGBA{0, 0, 255, 255}
	green = color.RGBA{0, 255, 0, 255}
)


func main() {
	display = m5tiny.InitDisplay()
	width, height := display.Size()

	display.FillRectangle(width/4, height/4, width/2, height/2, black)
	display.FilledCircle(width/2, height/2, 30, blue)
	display.WriteLine(&freemono.Regular9pt7b, 30, 40, "Hello M5Stack ...", green)
}
