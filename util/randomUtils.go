package util

import (
	"math/rand"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyz1234567890"

func RandString(length int) string {
	b := make([]byte, length)
	for i := range b {
		r := rand.Intn(len(letterBytes))
		b[i] = letterBytes[r]
	}
	return string(b)
}
