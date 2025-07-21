package utils

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"syscall"

	"golang.org/x/term"
)

const (
	chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

// GetRandomString generates a non-secure random string of a given length.
func GetRandomString(length int) string {
	result := make([]byte, length)
	for i := range length {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

// Prompt asks the user for input, with an option to hide the text for secrets.
func Prompt(msg string, secret bool) (string, error) {
	fmt.Print(msg)

	if secret {
		bytePassword, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			return "", err
		}
		fmt.Println()
		return strings.TrimSpace(string(bytePassword)), nil
	}

	reader := bufio.NewReader(os.Stdin)
	answer, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(answer), nil
}
