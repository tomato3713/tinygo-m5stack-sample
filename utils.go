//go:build m5stack

package m5tiny

import (
	"fmt"
	"image/color"
	"io/fs"
	"machine"

	// display
	"tinygo.org/x/drivers/ili9341"
	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinyfont"
)


type DisplayDevice struct {
	*ili9341.Device
}

// "tinygo.org/x/drivers/examples/ili9341/initdisplay" を参考に実装
func InitDisplay() *DisplayDevice {
	machine.SPI2.Configure(machine.SPIConfig{
		SCK:       machine.SPI0_SCK_PIN,
		SDO:       machine.SPI0_SDO_PIN,
		SDI:       machine.SPI0_SDI_PIN,
		Frequency: 40e6,
	})

	// configure backlight
	backlight := machine.LCD_BL_PIN
	backlight.Configure(machine.PinConfig{machine.PinOutput})

	display := ili9341.NewSPI(
		machine.SPI2,
		machine.LCD_DC_PIN,
		machine.LCD_SS_PIN,
		machine.LCD_RST_PIN,
	)

	// configure display
	display.Configure(ili9341.Config{
		Width:            320,
		Height:           240,
		DisplayInversion: true,
	})

	backlight.High()

	display.SetRotation(ili9341.Rotation0Mirror)

	return &DisplayDevice{
		display,
	}
}

func (d *DisplayDevice) FilledCircle(x int16, y int16, r int16, color color.RGBA) {
	tinydraw.FilledCircle(d, x, y, r, color)
}

func (d *DisplayDevice) WriteLine(font tinyfont.Fonter, x int16, y int16, str string, color color.RGBA) {
	tinyfont.WriteLine(d, font, x, y, str, color)
}

func (d *DisplayDevice) DrawImage(r fs.File, x, y, w, h int) {
	unit := 3
	buf := make([]uint8, w*h*unit)
	if n, err := r.Read(buf[:]); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("read bytes: ", n)
	}

	ws, we := 0, int(w)
	hs, he := 0, int(h)
	for yy := hs; yy < he; yy++ {
		for xx := ws; xx < we; xx++ {
			d.SetPixel(int16(x+xx), int16(y+yy), color.RGBA{buf[(xx+yy*we)*unit+2], buf[(xx+yy*we)*unit+1], buf[(xx+yy*we)*unit], 0})
		}
	}
}
