package main

import (
	"fmt"
	"math/rand"
	"time"
)

// const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
// const letters = "abcdefghijklmnopqrstuvwxyz"
const letters = "abcdef"

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

func findMatch(userInput string, strLen int) string {
	start := time.Now()

	// while loop until match
	for {

		a := randStr(strLen)

		if a ==  userInput {
			fmt.Println("A match was found: ", a)
			break
		}
	}
	elapsed := time.Since(start)
    fmt.Printf("compare function took %s\n\n", elapsed)
    return userInput
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	var userInput string
	var strLen int
	
	fmt.Println("Enter the desired string len (e.g. abcabcab is length of 8): ")
	fmt.Scanln(&strLen)

	fmt.Println("From the " + letters + " pool, please enter a combination of letters you would like to match: ")
	fmt.Scanln(&userInput)

	findMatch(userInput, strLen)

	// go compare()
	// go compare()
	time.Sleep(250 * time.Millisecond)
	fmt.Println("main done...")
	
}
