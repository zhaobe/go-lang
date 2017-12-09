package main

import "fmt"

func main() {
	fmt.Printf("hello, world\n")

	// int and float typing
	var num int = 10
	var floatNum float64 = 1.35134
	fmt.Println("my int number: ", num)
	fmt.Println("my float number: ", floatNum)

	// randNum is short declared as type int
	autoTypedNum := 24
	fmt.Println("my short variable declaration num is: ", autoTypedNum)

	// mixed variable declaration using type inference
	var a, b = 1, "notNum"
	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Printf("a is of type: %T \n", a)
	fmt.Printf("b is of type: %T \n\n", b)

	// arithmetic operations
	fmt.Println("10 + 5 = ", 10+5)
	fmt.Println("10 - 5 = ", 10-5)
	fmt.Println("10 * 5 = ", 10*5)
	fmt.Println("10 / 5 = ", 10/5)
	fmt.Println("10 % 5 = ", 10%5)

	// simple string with concat
	var myString string = "my simple string"
	fmt.Println("The length of "+myString+" is:", len(myString))

	// boolean and precision float
	var isOverTen bool = true
	fmt.Printf("floatNum with precision of 2: %.2f \n", floatNum)
	fmt.Printf("Returning data type for floatNum: %T\n", floatNum)
	fmt.Printf("Printing a boolean for isOverTen: %t\n\n", isOverTen)

	// simple for loop
	fmt.Println("for loops, different syntax")
	i := 1
	for i <= 5 {
		fmt.Println("for loop without semicolons: ", i)
		i++
	}

	for j := 1; j <= 5; j++ {
		fmt.Println("for loop with semicolons: ", j)
	}

	// switch statements
	someNum := 7

	switch someNum {
	case 17:
		fmt.Println("Your number is lower than 17")
	case 5:
		fmt.Println("Your number is higher than 5")
	default:
		fmt.Println("Your number is 7")
	}

}
