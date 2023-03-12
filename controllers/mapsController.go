package controllers

import (
	"context"
	"net/http"
	"strconv"
	"time"
	"tracy-api/helper"
	"tracy-api/services"

	"github.com/gofiber/fiber/v2"
)

type mapsHandler struct {
	mapsService services.MapsService
}

func NewMapsHandler(mapsService services.MapsService) *mapsHandler {
	return &mapsHandler{mapsService}
}

func (h *mapsHandler) GetDirection(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	origin := c.Query("origin")
	destination := c.Query("destination")

	direction, err := h.mapsService.GetDirection(ctx, origin, destination)

	if err != nil{
		response := helper.APIResponse("Failed to get direction", http.StatusBadRequest, "error", &fiber.Map{"error" : err})
		c.Status(http.StatusBadRequest).JSON(response)
		return nil
	}

	response := helper.APIResponse("Get direction success", http.StatusOK, "success", direction)
	c.Status(http.StatusOK).JSON(response)
	return nil
}

func (h *mapsHandler) GetGeocode(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	locationName := c.Query("location")

	geocode, err := h.mapsService.GetGeocode(ctx, locationName)

	if err != nil{
		response := helper.APIResponse("Failed to get geocode", http.StatusBadRequest, "error", &fiber.Map{"error" : err})
		c.Status(http.StatusBadRequest).JSON(response)
		return nil
	}

	response := helper.APIResponse("Get geocode success", http.StatusOK, "success", geocode)
	c.Status(http.StatusOK).JSON(response)
	return nil
}

func (h *mapsHandler) GetPoliceNearby(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	lat, _ := strconv.ParseFloat(c.Query("lat"), 64)
	
	lng, _ := strconv.ParseFloat(c.Query("lng"), 64)

	police, err := h.mapsService.GetPoliceNearby(ctx,lat,lng)

	if err != nil{
		response := helper.APIResponse("Failed to get police nearby", http.StatusBadRequest, "error", &fiber.Map{"error" : err})
		c.Status(http.StatusBadRequest).JSON(response)
		return nil
	}

	response := helper.APIResponse("Get police nearby success", http.StatusOK, "success", police)
	c.Status(http.StatusOK).JSON(response)
	return nil
}