package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  // only need to seed once
  rand.Seed(time.Now().UTC().UnixNano())
  var a int = randInt(50, 100)
  var b int = randInt(50, 100)
  fmt.Println(maxNum(a, b))
}

func randInt(min int, max int) int {
  return min + rand.Intn(max - min)
}

func maxNum(n1 int, n2 int) int {
  var res int
  var resStr string
  
  if (n1 > n2) {
    res = n1
    resStr = "A is greater than B"
  } else {
    res = n2
    resStr = "B is greater than A"
  }
  
  return res
}