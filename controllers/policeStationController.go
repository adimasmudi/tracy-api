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