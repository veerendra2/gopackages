package utils

import (
	"math/rand"
)

const (
	chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func RandomString(length int) string {
	result := make([]byte, length)
	for i := range length {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}
