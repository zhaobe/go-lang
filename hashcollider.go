package main

import (
  "fmt"
  "crypto/rand"
  "encoding/base64"
)

func randBytes(n int) ([]byte) {
  b := make([]byte, n)
  return b
}


func main() {
  fmt.Println(randBytes)
}