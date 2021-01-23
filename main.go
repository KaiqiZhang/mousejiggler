package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	go loop()
	fmt.Println("Press the Enter key to stop.")
	fmt.Scanln()
}

func loop() {
	for {
		nextX := rand.Intn(600) + 200
		nextY := rand.Intn(600) + 200
		robotgo.MoveSmooth(nextX, nextY)
		delay := 6 + rand.Intn(5)
		time.Sleep(time.Duration(delay) * time.Second)
	}
}
