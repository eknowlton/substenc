package main

import (
	"fmt"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"
const offset = 12

// Encrypt f func(plaintext string) string
func Encrypt(plaintext string) string {
	var result strings.Builder
	for i := 0; i < len(plaintext); i++ {
		if string(plaintext[i]) == " " {
			result.WriteString(" ")
			continue
		}

		currentPosition := strings.Index(alphabet, string(plaintext[i]))
		transposition := currentPosition + offset
		if transposition > len(alphabet) {
			transposition = transposition - len(alphabet)
		}

		result.WriteString(string(alphabet[transposition]))
	}
	return result.String()
}

// Decrypt f func(cipertext string) string
func Decrypt(cipertext string) string {
	var result strings.Builder
	for i := 0; i < len(cipertext); i++ {
		if string(cipertext[i]) == " " {
			result.WriteString(" ")
			continue
		}

		currentPosition := strings.Index(alphabet, string(cipertext[i]))
		transposition := currentPosition - offset
		if transposition < 0 {
			transposition = transposition + len(alphabet)
		}

		result.WriteString(string(alphabet[transposition]))
	}
	return result.String()
}

func main() {
	fmt.Println("Encrypt string 'Hello World 123': ")
	result := Encrypt("Hello World 123")
	fmt.Printf("Decrypt result: '%s': \n", result)
	fmt.Println(Decrypt(result))
}
