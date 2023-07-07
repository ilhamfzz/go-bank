package util

import (
	"math/rand"
	"time"
)

func GenerateRandomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	generate := make([]byte, n)
	for i := range generate {
		generate[i] = letters[rand.Intn(len(letters))]
	}

	return string(generate)
}

func GenerateRandomOTPNumber(n int) string {
	const numbers = "0123456789"
	rand.Seed(time.Now().UnixNano())
	generate := make([]byte, n)
	for i := range generate {
		generate[i] = numbers[rand.Intn(len(numbers))]
	}

	return string(generate)
}
