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

func main() {
  rand.Seed(time.Now().UTC().UnixNano())
  for i := 0; i < 15; i++ {
    fmt.Println("t1: ", randStr(8), "\tt2: ", randStr(8))
  }
  
}