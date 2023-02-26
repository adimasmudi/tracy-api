package controllers

import (
	"context"
	"net/http"
	"time"
	"tracy-api/helper"
	"tracy-api/inputs"
	"tracy-api/services"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type reportHandler struct {
	reportService services.ReportService
}

func NewReportHandler(reportService services.ReportService) *reportHandler{
	return &reportHandler{reportService}
}

func (h *reportHandler) CreateReport(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var input inputs.CreateReportInput
	
	//validate the request body
	if err := c.BodyParser(&input); err != nil {
		response := helper.APIResponse("Report Failed, wrong input format", http.StatusBadRequest, "error", &fiber.Map{"error" : err})
		c.Status(http.StatusBadRequest).JSON(response)
		return nil
	}

	currentEmailUser := c.Locals("currentUserEmail").(string)

	reported, err := h.reportService.CreateReport(ctx,currentEmailUser, input)

	if err != nil{
		response := helper.APIResponse("Report Failed", http.StatusBadRequest, "error", &fiber.Map{"error" : err})
		c.Status(http.StatusBadRequest).JSON(response)
		return nil
	}

	response := helper.APIResponse("Create report success", http.StatusOK, "success", reported)
	c.Status(http.StatusOK).JSON(response)
	return nil
}

func (h *reportHandler) GetDetailReportById(c *fiber.Ctx) error{
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id :=c.Params("id")

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil{
		response := helper.APIResponse("Failed to get report detail", http.StatusBadRequest, "error", &fiber.Map{"error" : err})
		c.Status(http.StatusBadRequest).JSON(response)
		return nil
	}

	report, err := h.reportService.GetById(ctx, objectId)

	if err != nil{
		response := helper.APIResponse("Failed to get report detail", http.StatusBadRequest, "error", &fiber.Map{"error" : err})
		c.Status(http.StatusBadRequest).JSON(response)
		return nil
	}

	response := helper.APIResponse("get report detail data success", http.StatusOK, "success", report)
	c.Status(http.StatusOK).JSON(response)
	return nil
}
