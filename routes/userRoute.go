package routes

import (
	"tracy-api/controllers"
	"tracy-api/repository"
	"tracy-api/services"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func UserRoute(api fiber.Router, collection *mongo.Collection) {
	
	userRepository := repository.NewUserRepository(collection)
	userService := services.NewUserService(userRepository)
	userHandler := controllers.NewUserHandler(userService)

	authUser := api.Group("/auth/users")

	authUser.Post("/login",userHandler.Login)
		
}