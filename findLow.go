package main

import (
	"fmt"
)

func main() {
	// find the lowest number in the array
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	low := 0
	for i := 0; i < len(x); i++ {
		for j := i + 1; j < len(x); j++ {
			if x[i] <= x[j] {
				low = x[i]
			} else if x[j] <= x[i] {
				low = x[j]
			}
		}
	}
	fmt.Println(low)
}
