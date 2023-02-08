package helper

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}