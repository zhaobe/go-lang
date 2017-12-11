package main

import (
	"fmt"
	"math/rand"
	"time"
)

// const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const letters = "abc"

func randStr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func compStr(a, b string) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return +1
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	a := randStr(3)
	var b string
	
	for {
			b := randStr(3)
			fmt.Println("a: ", a, "\tb: ", b)
			fmt.Println("compare: ", compStr(a, b))
			if a == b {
				break
			}
		}
}
