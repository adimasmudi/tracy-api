package controllers

import (
	"context"
	"time"
	"tracy-api/services"

	"github.com/gofiber/fiber"
)

type userHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) Login(c *fiber.Ctx) {
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	return
}