package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
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
	decryptPtr := flag.Bool("decrypt", false, "Decrypt input")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		if *decryptPtr {
			fmt.Println(Decrypt(scanner.Text()))
		} else {
			fmt.Println(Encrypt(scanner.Text()))
		}
	}
}
