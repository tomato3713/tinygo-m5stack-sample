package main

import (
	"embed"
	"github.com/tomato3713/m5tiny"
	"fmt"
)

//go:embed img/*
var static embed.FS

var (
	display *m5tiny.DisplayDevice
)

func main() {
	display = m5tiny.InitDisplay()

	image := "img/image01.bin"
	r, err := static.Open(image)
	if err != nil {
		fmt.Println("file open err: ", err)
	}
	display.DrawImage(r, 60, 20, 200, 200)
}
