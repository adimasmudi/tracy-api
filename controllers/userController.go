package controllers

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"tracy-api/configs"
	"tracy-api/helper"
	"tracy-api/services"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
)

type userHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) Callback(c *fiber.Ctx){
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if c.FormValue("state") != os.Getenv("oAuth_String") {
		response := helper.APIResponse("Can't login to your account", http.StatusBadRequest, "error", nil)
		c.Status(http.StatusBadRequest).JSON(response)
		return
	}

	token, err := configs.GoogleOAuthConfig().Exchange(oauth2.NoContext, c.FormValue("code"))
	if err != nil {
		response := helper.APIResponse("code exchange failed", http.StatusBadRequest, "error", err)
		c.Status(http.StatusBadRequest).JSON(response)
		return
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		response := helper.APIResponse("failed getting user info", http.StatusBadRequest, "error", err)
		c.Status(http.StatusBadRequest).JSON(response)
		return
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		response := helper.APIResponse("Failed reading response body", http.StatusBadRequest, "error", err)
		c.Status(http.StatusBadRequest).JSON(response)
		return
	}

	var googleUser helper.GoogleUser

	json.Unmarshal([]byte(string(contents)), &googleUser)

	user, err := h.userService.Signup(googleUser)

	if err != nil{
		response := helper.APIResponse("Login User Failed", http.StatusBadRequest, "error", err)
		c.Status(http.StatusBadRequest).JSON(response)
		return
	}
	
	c.Status(http.StatusBadRequest).JSON(user)
}