package utils

import (
	"healthcare-capt-america/enums"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(str string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(str), enums.DefaultCost)
	return string(b), err
}
