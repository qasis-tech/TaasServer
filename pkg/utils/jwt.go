package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Username string `json:"username"`
	UserId   uint   `json:"userid"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateJwtToken(username string, userid uint) (string, error) {
	Claims := &Claims{
		Username: username,
		UserId:   userid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (any, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil

}
