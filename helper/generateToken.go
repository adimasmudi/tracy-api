package helper

import (
	"os"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(email string) (string, error) {
	claim := jwt.MapClaims{}
	claim["email"] = email
	claim["isExpired"] = false 

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}