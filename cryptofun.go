package main

import (
  "fmt"
  "bufio"
  "os"
//  "crypto/aes"
//  "crypto/cipher"
//  "crypto/rand"
//  "crypto/base64"
)

func main() {
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Enter text: ")
  text, _ := reader.ReadString('\n')
  fmt.Println(text)
}
