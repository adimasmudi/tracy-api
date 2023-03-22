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
	"tracy-api/inputs"
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

	// Create cookie
	// cookie := new(fiber.Cookie)
	// cookie.Name = "user"
	// cookie.Value = user.Email
	// cookie.Expires = time.Now().Add(24 * time.Hour)

	// fmt.Println(cookie)
  
	// // Set cookie
	// c.Cookie(cookie)

	formatter := formatter.FormatUser(user)
	responses := helper.APIResponse("Signup User Success", http.StatusOK, "success", &fiber.Map{"user" : formatter, "token" : loginToken})
	c.Status(http.StatusOK).JSON(responses)
	return nil
}

func (h *userHandler) Register(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var input inputs.RegisterUserInput

	// file, err := c.FormFile("picture")

	// if err != nil{
	// 	response := helper.APIResponse("Wrong file format", http.StatusBadRequest, "error", &fiber.Map{"error" : err})
	// 	c.Status(http.StatusBadRequest).JSON(response)
	// 	return nil
	// }

	//validate the request body
	if err := c.BodyParser(&input); err != nil {
		response := helper.APIResponse("Register Failed", http.StatusBadRequest, "error", &fiber.Map{"error" : err})
		c.Status(http.StatusBadRequest).JSON(response)
		return nil
	}

	// fileToSave := fmt.Sprintf("./images/%s-%s",input.Username,file.Filename)

	registeredUser, err := h.userService.Register(ctx, input)

	if err != nil{
		response := helper.APIResponse("Register Failed", http.StatusBadRequest, "error", &fiber.Map{"error" : err})
		c.Status(http.StatusBadRequest).JSON(response)
		return nil
	}

	// c.SaveFile(file, fileToSave)

	response := helper.APIResponse("User register success", http.StatusOK, "success", registeredUser)
	c.Status(http.StatusOK).JSON(response)
	return nil
}

func (h *userHandler) Login(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var input inputs.LoginUserInput

	//validate the request body
	if err := c.BodyParser(&input); err != nil {
		response := helper.APIResponse("Login Failed", http.StatusBadRequest, "error", &fiber.Map{"error" : err})
		c.Status(http.StatusBadRequest).JSON(response)
		return nil
	}

	logedinUser, token,  err := h.userService.Login(ctx,input)

	if err != nil{
		response := helper.APIResponse("Login Failed", http.StatusBadRequest, "error", &fiber.Map{"error" : err})
		c.Status(http.StatusBadRequest).JSON(response)
		return nil
	}

	response := helper.APIResponse("Login success", http.StatusOK, "success", &fiber.Map{"user" : logedinUser, "token" : token})
	c.Status(http.StatusOK).JSON(response)
	return nil
}

func (h *userHandler) GetProfile(c *fiber.Ctx)error{
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	currentEmailUser := c.Locals("currentUserEmail").(string)

	user, err := h.userService.GetProfile(ctx,currentEmailUser)

	if err != nil{
		response := helper.APIResponse("Can't get user data", http.StatusBadRequest, "error", &fiber.Map{"error" : err})
		c.Status(http.StatusBadRequest).JSON(response)
		return nil
	}

	response := helper.APIResponse("Get user data success", http.StatusOK, "success", user)
	c.Status(http.StatusOK).JSON(response)
	return nil
}

func (h *userHandler) UpdateProfile(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	currentEmailUser := c.Locals("currentUserEmail").(string)

	var input inputs.UpdateUserInput

	//validate the request body
	if err := c.BodyParser(&input); err != nil {
		response := helper.APIResponse("Update Profile Failed", http.StatusBadRequest, "error", &fiber.Map{"error" : err})
		c.Status(http.StatusBadRequest).JSON(response)
		return nil
	}

	updatedUser, err := h.userService.UpdateProfile(ctx, currentEmailUser, input)

	if err != nil{
		response := helper.APIResponse("Can't update user data", http.StatusBadRequest, "error", &fiber.Map{"error" : err})
		c.Status(http.StatusBadRequest).JSON(response)
		return nil
	}

	response := helper.APIResponse("Update user data success", http.StatusOK, "success", updatedUser)
	c.Status(http.StatusOK).JSON(response)
	return nil

}

// func (h *userHandler) Logout(c *fiber.Ctx) error{
// 	return nil
// }