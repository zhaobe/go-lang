package main

import (
	"fmt"
	"math/rand"
	"time"
)

// const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
// const letters = "abcdefghijklmnopqrstuvwxyz"
const letters = "abc"

func randStr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func compare() {
	start := time.Now()

	a := randStr(8)
	// while loop until match
	for {

		b := randStr(8)

		if a == b {
			fmt.Println("a: ", a, "\tb: ", b)
			break
		}
	}
	elapsed := time.Since(start)
    fmt.Printf("compare function took %s\n\n", elapsed)
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	go compare()
	go compare()
	time.Sleep(250 * time.Millisecond)
	fmt.Println("main done...")
	
}
