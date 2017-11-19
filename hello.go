package main

import "fmt"

func main() {
  fmt.Printf("hello, world\n")
  
  // int and float typing
  var num int = 10
  var floatNum float64 = 1.35134
  fmt.Println("my int number: ", num)
  fmt.Println("my float number: ", floatNum)
  
  // randNum is autoset to type int here
  autoTypedNum := 24
  fmt.Println("my auto typed num: ", autoTypedNum)
  
  // arithmetic operations
  fmt.Println("10 + 5 = ", 10 + 5)
  fmt.Println("10 - 5 = ", 10 - 5)
  fmt.Println("10 * 5 = ", 10 * 5)
  fmt.Println("10 / 5 = ", 10 / 5)
  fmt.Println("10 % 5 = ", 10 % 5)
  
}

