package main

import (
	"fmt"
)

func main() {

	lineOne := 10
	lineTwo := [10]int{64630, 11735, 14216, 99233, 14470, 4978, 73429, 38120, 51135, 67060}

	var sum int
	for i := 0; i < lineOne; i++ {
		sum += lineTwo[i]
	}
	mean := sum/lineOne
	fmt.Println(mean)

	
}