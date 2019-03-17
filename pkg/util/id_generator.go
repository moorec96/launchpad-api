package util

import (
	"crypto/rand"
)

func GenerateRandomString(length int) string {
	b := make([]byte, length)
	rand.Read(b)
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	for i, a := range b {
		b[i] = letters[a%byte(len(letters))]
	}
	return string(b)
}
