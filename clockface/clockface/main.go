package main

import (
	"hello/clockface/svg"
	"os"
	"time"
)

func main() {
	t := time.Now()

	file, err := os.Create("clock.svg")
	if err != nil {
		// Handle error
	}
	defer file.Close()
	svg.Write(file, t)
}
