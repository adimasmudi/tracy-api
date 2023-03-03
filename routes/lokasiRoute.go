package routes

import (
	"tracy-api/controllers"
	"tracy-api/middlewares"
	"tracy-api/repository"
	"tracy-api/services"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func LokasiRoute(api fiber.Router, lokasiCollection *mongo.Collection) {
	lokasiRepository := repository.NewLokasiRepository(lokasiCollection)
	lokasiService := services.NewLokasiService(lokasiRepository)
	lokasiHandler := controllers.NewLokasiHandler(lokasiService)

	lokasiAPI := api.Group("/lokasi")
	lokasiAPI.Post("/save",middlewares.Auth,lokasiHandler.SaveLocation)
}