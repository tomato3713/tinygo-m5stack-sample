package main

import (
	"fmt"
	"time"
)

func main() {
	count := 0
	for {
		fmt.Println(count, ": Hello, World")
		time.Sleep(time.Millisecond * 1000)
		count++
	}
}
