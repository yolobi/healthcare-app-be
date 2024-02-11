package security

import (
	"healthcare-capt-america/entities/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JWT_KEY = []byte(os.Getenv("JWT_SECRET"))

type JWTClaim struct {
	UserID uint64
	Email  string
	Role   string
	jwt.RegisteredClaims
}

func GenerateJWT(user *models.User, role string) string {
	expTime := time.Now().Add(time.Hour * 6)
	claims := JWTClaim{
		UserID: user.ID,
		Email:  user.Email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "APPLICATION",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := tokenAlgo.SignedString(JWT_KEY)
	return token
}
