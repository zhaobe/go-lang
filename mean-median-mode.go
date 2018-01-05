package main

import (
	"fmt"
	"sort"
)

func main() {

	lineOne := 10
	lineTwo := []int{64630, 11735, 14216, 99233, 14470, 4978, 73429, 38120, 51135, 67060}

	var sum int
	for i := 0; i < lineOne; i++ {
		sum += lineTwo[i]
	}
	mean := sum/lineOne
	fmt.Println("Mean:\t", mean)

	if(len(lineTwo) % 2 == 0) {
		midOne := (len(lineTwo) / 2) - 1
		midTwo := len(lineTwo) / 2
		
		sort.Ints(lineTwo)
		median := (lineTwo[midOne] + lineTwo[midTwo]) / 2
		fmt.Println("Median:\t", median)
		fmt.Println("even")
	} else {
		midPoint := (len(lineTwo) / 2) + 1
		median := lineTwo[midPoint]
		fmt.Println("Median:\t", median)
	}

	


}