package middlewares

import (
	"errors"
	"net/http"
	"os"
	"strings"
	"tracy-api/helper"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Auth(c *fiber.Ctx)error{
	authorizationHeader := c.Get("Authorization")
	
	if !strings.Contains(authorizationHeader, "Bearer"){
		response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
		c.Status(http.StatusUnauthorized).JSON(response)
		return nil
	}

	tokenArray := strings.Split(authorizationHeader," ")
	if len(tokenArray) < 2{
		response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
		c.Status(http.StatusUnauthorized).JSON(response)
		return nil
	}

	tokenString := tokenArray[1]

	token, err := validateToken(tokenString)

	if err != nil{
		response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
		c.Status(http.StatusUnauthorized).JSON(response)
		return err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid{
		response := helper.APIResponse("Unauthorized, token is not valid", http.StatusUnauthorized, "error", nil)
		c.Status(http.StatusUnauthorized).JSON(response)
		return nil
	}

	email := claims["email"].(string)

	c.Locals("currentUserEmail",email)


	return c.Next()
}

func validateToken(encodedToken string)(*jwt.Token, error){
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token)(interface{}, error){
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok{
			return nil, errors.New("invalid token")
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil{
		return token, err
	}

	return token, nil
}