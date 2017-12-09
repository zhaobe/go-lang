package main

import (
	"fmt"
	"math/rand"
	"time"
)

func loop(num int) {
	for i := 0; i < 15; i++ {
		fmt.Println(num, "\t", i)
		randNum := time.Duration(rand.Intn(500))
		time.Sleep(time.Millisecond * randNum)
	}
}

func main() {
	for i := 0; i < 15; i++ {
		go loop(i)
	}
	var userInput string
	fmt.Scanln(&userInput)
}
