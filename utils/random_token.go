package utils

import (
	"crypto/rand"
	"fmt"
)

func GenerateToken(len int) string {
	b := make([]byte, len)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
