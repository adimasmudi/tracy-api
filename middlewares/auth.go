package middlewares

import (
	"errors"
	"net/http"
	"strings"
	"tracy-api/helper"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Auth(c *fiber.Ctx)error{
	
	authorizationHeader := c.Get("Authorization")
	
	if !strings.Contains(authorizationHeader, "Bearer"){
		response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", errors.New("you have to use bearer"))
		c.Status(http.StatusUnauthorized).JSON(response)
		return nil
	}

	tokenArray := strings.Split(authorizationHeader," ")
	if len(tokenArray) < 2{
		response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", errors.New("can't Find token"))
		c.Status(http.StatusUnauthorized).JSON(response)
		return nil
	}

	tokenString := tokenArray[1]

	token, err := helper.ValidateToken(tokenString)

	if err != nil{
		response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", errors.New("token is not valid"))
		c.Status(http.StatusUnauthorized).JSON(response)
		return err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid{
		response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", errors.New("token is not valid"))
		c.Status(http.StatusUnauthorized).JSON(response)
		return nil
	}

	email := claims["email"].(string)

	// cookie := c.Cookies("email","") // set default to empty

	// if cookie == ""{
	// 	response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", errors.New("you have to logged in again"))
	// 	c.Status(http.StatusUnauthorized).JSON(response)
	// 	return nil
	// }
	
	c.Locals("currentUserEmail",email)

	return c.Next()
}

