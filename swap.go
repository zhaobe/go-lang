package main

import "fmt"

// swapping without temp variable
func main() {
	var a int = 3
	var b int = 6

	fmt.Printf("Before swap: a = %d, b = %d \n", a, b)
	a = a + b
	b = a - b
	a = a - b

	fmt.Printf("After swap: a = %d, b = %d \n", a, b)
}
