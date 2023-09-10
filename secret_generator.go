package generator

import "math/rand"

const (
	_ASCII_START = 48
	_ASCII_END   = 122
)

func GenerateSecret(length int) string {
	secret := ""
	for i := 0; i < length; i++ {
		char := rand.Intn(_ASCII_END-_ASCII_START) + _ASCII_START
		secret = secret + string(rune(char))
	}
	return secret
}
