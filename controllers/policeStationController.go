package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"tracy-api/helper"
	"tracy-api/inputs"
	"tracy-api/services"

	"github.com/gofiber/fiber/v2"
)

type policeStationHandler struct {
	policeStationService services.PoliceStationService
}

func NewPoliceStationHandler(policeStationService services.PoliceStationService) *policeStationHandler{
	return &policeStationHandler{policeStationService}
}

func (h *policeStationHandler) Register(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var input inputs.PoliceStationInput

	file, err := c.FormFile("picture")

	if err != nil{
		response := helper.APIResponse("Wrong file format", http.StatusBadRequest, "error", &fiber.Map{"error" : err})
		c.Status(http.StatusBadRequest).JSON(response)
		return nil
	}

	//validate the request body
	if err := c.BodyParser(&input); err != nil {
		response := helper.APIResponse("Register Failed", http.StatusBadRequest, "error", &fiber.Map{"error" : err})
		c.Status(http.StatusBadRequest).JSON(response)
		return nil
	}

	fileToSave := fmt.Sprintf("./images/%s-%s",input.Username,file.Filename)

	registeredUser, err := h.policeStationService.Save(ctx, input, fileToSave)

	if err != nil{
		response := helper.APIResponse("Register Failed", http.StatusBadRequest, "error", &fiber.Map{"error" : err})
		c.Status(http.StatusBadRequest).JSON(response)
		return nil
	}

	c.SaveFile(file, fileToSave)

	response := helper.APIResponse("Police station register success", http.StatusOK, "success", registeredUser)
	c.Status(http.StatusOK).JSON(response)
	return nil
}

func (h *policeStationHandler) Login(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var input inputs.PoliceStationLoginInput

	//validate the request body
	if err := c.BodyParser(&input); err != nil {
		response := helper.APIResponse("Login Failed", http.StatusBadRequest, "error", &fiber.Map{"error" : err})
		c.Status(http.StatusBadRequest).JSON(response)
		return nil
	}

	logedinUser, token,  err := h.policeStationService.Login(ctx,input)

	if err != nil{
		response := helper.APIResponse("Login Failed", http.StatusBadRequest, "error", &fiber.Map{"error" : err})
		c.Status(http.StatusBadRequest).JSON(response)
		return nil
	}

	// Create cookie
	cookie := new(fiber.Cookie)
	cookie.Name = "email"
	cookie.Value = logedinUser.Email
	cookie.Expires = time.Now().Add(24 * time.Hour)

	fmt.Println(cookie)
  
	// // Set cookie
	c.Cookie(cookie)

	response := helper.APIResponse("Login success", http.StatusOK, "success", &fiber.Map{"police" : logedinUser, "token" : token})
	c.Status(http.StatusOK).JSON(response)
	return nil
}

func (h *policeStationHandler) GetProfile(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	currentEmailPoliceDashboard := c.Locals("currentUserEmail").(string)

	police, err := h.policeStationService.GetProfile(ctx, currentEmailPoliceDashboard)

	if err != nil{
		response := helper.APIResponse("Can't get police station data", http.StatusBadRequest, "error", &fiber.Map{"error" : err})
		c.Status(http.StatusBadRequest).JSON(response)
		return nil
	}

	response := helper.APIResponse("get police station data success", http.StatusOK, "success", police)
	c.Status(http.StatusOK).JSON(response)
	return nil
}

func (h *policeStationHandler) Logout(c *fiber.Ctx) error{
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cookie := new(fiber.Cookie)
	cookie.Name = "email"
	cookie.Value = ""
	cookie.Expires = time.Now().Add(-time.Hour * 24)
	c.Cookie(cookie)
	
	response := helper.APIResponse("logout police station success", http.StatusOK, "success", nil)
	c.Status(http.StatusOK).JSON(response)
	return nil
}