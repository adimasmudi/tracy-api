package routes

import (
	"tracy-api/controllers"
	"tracy-api/middlewares"
	"tracy-api/repository"
	"tracy-api/services"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func ReportRoute(api fiber.Router, collection []*mongo.Collection){

	reportRepository := repository.NewReportRepository(collection[0])
	userRepository := repository.NewUserRepository(collection[1])
	policeStationRepository := repository.NewPoliceStationRepository(collection[2])

	reportService := services.NewReportService(reportRepository, userRepository, policeStationRepository)
	reportHandler := controllers.NewReportHandler(reportService)

	reportAPI := api.Group("/report")
	reportAPI.Post("/create",middlewares.Auth, reportHandler.CreateReport)
	reportAPI.Get("/detail/:id", middlewares.Auth, reportHandler.GetDetailReportById)
	reportAPI.Get("/all", middlewares.Auth, reportHandler.GetAllReport)
	reportAPI.Get("/current/all", middlewares.Auth, reportHandler.GetAllByCurrentUser)
	reportAPI.Put("/updateStatus/:id", middlewares.Auth, reportHandler.UpdateStatus)
}