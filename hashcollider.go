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

// This function generates one hash for A and then tries to find a match for B
func compare(strLen int) {
	start := time.Now()

	a := randStr(strLen)
	var counter int = 0
	
	// while loop until match
	for {
		counter++
		b := randStr(strLen)
		// fmt.Println("a: ", a, "\tb: ", b)

		if a == b {
			fmt.Println("This was the matching sequence \ta: ", a, "\tb: ", b)
			break
		}
	}
	elapsed := time.Since(start)
    fmt.Printf("compare function took %s\n\n", elapsed)
    fmt.Printf("and it generated %d hashes to find this match\n", counter)
}

// This function requires user input
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
	time.Sleep(250 * time.Millisecond)
    fmt.Printf("compare function took %s\n\n", elapsed)
    return userInput
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	// var userInput string
	// var strLen int
	
	// fmt.Println("Enter the desired string len (e.g. abcabcab is length of 8): ")
	// fmt.Scanln(&strLen)

	// fmt.Println("From the " + letters + " pool, please enter a combination of letters you would like to match: ")
	// fmt.Scanln(&userInput)

	// go findMatch(userInput, strLen)
	
	compare(3)

	// go compare()
	// go compare()
	time.Sleep(250 * time.Millisecond)
	fmt.Println("main done...")
	
}
