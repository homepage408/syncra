package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	SecretKey string
}

func (j *JWT) GenerateToken(userID string) (string, error) {

	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(j.SecretKey))
}
