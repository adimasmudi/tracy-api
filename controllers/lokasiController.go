package controllers

import (
	"context"
	"net/http"
	"time"
	"tracy-api/helper"
	"tracy-api/inputs"
	"tracy-api/services"

	"github.com/gofiber/fiber/v2"
)

type lokasiHandler struct {
	lokasiService services.LokasiService
}

func NewLokasiHandler(lokasiService services.LokasiService) *lokasiHandler{
	return &lokasiHandler{lokasiService}
}

func (h *lokasiHandler) SaveLocation(c *fiber.Ctx) error{
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var input inputs.AddLokasiInput

	//validate the request body
	if err := c.BodyParser(&input); err != nil {
		response := helper.APIResponse("Report Failed, wrong input format", http.StatusBadRequest, "error", &fiber.Map{"error" : err})
		c.Status(http.StatusBadRequest).JSON(response)
		return nil
	}

	savedLocation, err := h.lokasiService.SaveLocation(ctx, input)

	if err != nil{
		response := helper.APIResponse("Save Location Failed", http.StatusBadRequest, "error", &fiber.Map{"error" : err})
		c.Status(http.StatusBadRequest).JSON(response)
		return nil
	}

	response := helper.APIResponse("Save Location success", http.StatusOK, "success", savedLocation)
	c.Status(http.StatusOK).JSON(response)
	return nil

}