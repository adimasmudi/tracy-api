package routes

import (
	"tracy-api/controllers"
	"tracy-api/repository"
	"tracy-api/services"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func PoliceStationRoute(api fiber.Router, policeStationCollection *mongo.Collection){
	
	policeStationRepository := repository.NewPoliceStationRepository(policeStationCollection)
	policeStationService := services.NewPoliceStationService(policeStationRepository)
	policeStationHandler := controllers.NewPoliceStationHandler(policeStationService)

	policeAPI := api.Group("/police")
	policeAPI.Post("/register", policeStationHandler.Register)
}