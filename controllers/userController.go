package controllers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"
	"tracy-api/configs"
	formatter "tracy-api/formatters"
	"tracy-api/helper"
	"tracy-api/services"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) Callback(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if c.FormValue("state") != os.Getenv("oAuth_String") {
		response := helper.APIResponse("Can't login to your account", http.StatusBadRequest, "error", nil)
		c.Status(http.StatusBadRequest).JSON(response)
		return nil
	}

	token, err := configs.GoogleOAuthConfig().Exchange(context.Background(), c.FormValue("code"))
	if err != nil {
		response := helper.APIResponse("code exchange failed", http.StatusBadRequest, "error", fiber.Map{"error" : err})
		c.Status(http.StatusBadRequest).JSON(response)
		return nil
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		response := helper.APIResponse("failed getting user info", http.StatusBadRequest, "error", fiber.Map{"error" : err})
		c.Status(http.StatusBadRequest).JSON(response)
		return nil
	}

	defer response.Body.Close()
	contents, err := io.ReadAll(response.Body)
	if err != nil {
		response := helper.APIResponse("Failed reading response body", http.StatusBadRequest, "error", fiber.Map{"error" : err})
		c.Status(http.StatusBadRequest).JSON(response)
		return nil
	}

	var googleUser helper.GoogleUser

	json.Unmarshal([]byte(string(contents)), &googleUser)

	user,loginToken, err := h.userService.Signup(ctx,googleUser)

	if err != nil{
		response := helper.APIResponse("Signup User Failed", http.StatusBadRequest, "error", &fiber.Map{"error" : err})
		c.Status(http.StatusBadRequest).JSON(response)
		return nil
	}
	formatter := formatter.FormatUser(user)
	responses := helper.APIResponse("Signup User Success", http.StatusOK, "success", &fiber.Map{"user" : formatter, "token" : loginToken})
	c.Status(http.StatusOK).JSON(responses)
	return nil
}