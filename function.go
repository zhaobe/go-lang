package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// only need to seed once
	rand.Seed(time.Now().UTC().UnixNano())
	var a int = randInt(50, 100)
	var b int = randInt(50, 100)

	maxNum(a, b)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func maxNum(n1 int, n2 int) {
	if n1 > n2 {
		fmt.Printf("(A: %d) is greater than (B: %d)\n", n1, n2)
	} else {
		fmt.Printf("(B: %d) is greater than (A: %d)\n", n2, n1)
	}
}
