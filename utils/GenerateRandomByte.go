package utils

import (
	"crypto/rand"
)

func GenerateRandomBytes() []byte {
	randomBytes := make([]byte, 6)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil
	}
	return randomBytes
}
