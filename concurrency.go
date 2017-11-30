package main

import "fmt"

func loop(num int) {
  for i := 0; i < 100; i++ {
    fmt.Println(num, "", i)
  }
}

func main() {
  go loop(0)
  var userInput string
  fmt.Scanln(&userInput)
}