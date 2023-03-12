package routes

import (
	"tracy-api/controllers"
	"tracy-api/middlewares"
	"tracy-api/repository"
	"tracy-api/services"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

// convert location name to geocode

// search police nearby

func MapsRoute(api fiber.Router, policeStationCollection *mongo.Collection){
	
	policeStationRepository := repository.NewPoliceStationRepository(policeStationCollection)

	mapsService := services.NewMapsService(policeStationRepository)
	mapsHandler := controllers.NewMapsHandler(mapsService)

	mapsAPI := api.Group("/maps")
	mapsAPI.Get("/direction", middlewares.Auth, mapsHandler.GetDirection)
	mapsAPI.Get("/geocode", middlewares.Auth, mapsHandler.GetGeocode)
	mapsAPI.Get("/nearby", middlewares.Auth, mapsHandler.GetPoliceNearby)
}