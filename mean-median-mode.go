package main

import (
	"fmt"
	"sort"
)

func mean(lineOne int, lineTwo []int) {
	var sum int
	// the i < lineOne could also be len(lineTwo), but since lineOne exist, we will use that
	for i := 0; i < lineOne; i++ {
		sum += lineTwo[i]
	}
	mean := sum / lineOne
	fmt.Println("Mean:\t", mean)
}

func median(lineOne int, lineTwo []int) {
	if len(lineTwo)%2 == 0 {
		midOne := (len(lineTwo) / 2) - 1
		midTwo := len(lineTwo) / 2

		sort.Ints(lineTwo)
		median := (lineTwo[midOne] + lineTwo[midTwo]) / 2
		fmt.Println("Median:\t", median)
	} else {
		midPoint := (len(lineTwo) / 2) + 1
		median := lineTwo[midPoint]
		fmt.Println("Median:\t", median)
	}
}

func mode(lineOne int, lineTwo []int) {
	mode := 0
	maxCount := 0

	for i := 0; i < lineOne; i++ {
		count := 0
		for j := 0; j < lineOne; j++ {
			if lineTwo[j] == lineTwo[i] {
				count++
			}
		}
		if count > maxCount {
			maxCount = count
			mode = lineTwo[i]
		}
	}
	fmt.Println("Mode:\t", mode)

}

func main() {

	lineOne := 10
	lineTwo := []int{64630, 11735, 14216, 99233, 14470, 4978, 73429, 38120, 51135, 67060}
	mean(lineOne, lineTwo)
	median(lineOne, lineTwo)
	mode(lineOne, lineTwo)
}
