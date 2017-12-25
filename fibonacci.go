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

// tabulation, bottom up approach
func fib_tab(n int) int {
	var base = n + 1
	tab := make([]int, base)
	
	tab[0] = 0
	tab[1] = 1
	for i := 2; i <= n; i++ {
		tab[i] = tab[i - 1] + tab[i - 2]
	}
	return tab[n]
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
	fmt.Printf("The memoization Fibonacci number is: %d\n", fib(input))
	fmt.Printf("The tabulation Fibonacci number is: %d\n", fib_tab(input))
}