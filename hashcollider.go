package main

import (
	"fmt"
	"math/rand"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

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
	for i := 0; i < 15; i++ {
		a := randStr(8)
		b := randStr(8)
		fmt.Println("a: ", a, "\tb: ", b)
		fmt.Println("compare: ", compStr(a, b))
	}

}
