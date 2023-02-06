package main

import (
	"tracy-api/configs"
	"tracy-api/routes"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func main(){
	app := fiber.New()

	// run database
	configs.ConnectDB()

	// collections
	var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")

	api := app.Group("/api/v1")

	// routes
	routes.UserRoute(api, userCollection)

	app.Listen(":5000")
}