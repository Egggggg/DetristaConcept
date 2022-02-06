package utils

import "math/rand"

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func RandomString(length int) string {
	runes := make([]rune, length)

	for i := range runes {
		runes[i] = letters[rand.Intn(len(letters))]
	}

	return string(runes)
}
