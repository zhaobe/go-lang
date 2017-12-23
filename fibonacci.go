package main

import (
	"fmt"
)

const MAX = 100
const NIL = -1
var table[MAX] int

func memoize() {
	for i := 0; i < MAX; i++ {
		table[i] = NIL
	}
}

func fib(n int) int {
	if (table[n] == NIL) {
		if (n <= 1) {
			table[n] = n
		} else {
			table[n] = fib(n - 1) + fib(n - 2)
		}
	}
	
	return table[n]
}

func main() {
	var input int
	
	memoize()
	fmt.Println("Enter a number between 0 and 100: ")
	fmt.Scanln(&input)
	fmt.Printf("The Fibonacci number is: %d\n", fib(input))
}