package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println("Original text: ", text)

	key := []byte("thisisthekey")
	encText := encrypt(text, key)
	fmt.Println(encText)
}

func encrypt(text string, key []byte) string {
	plain := []byte(text)

}
