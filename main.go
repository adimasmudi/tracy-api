package main

import (
	"tracy-api/configs"

	"tracy-api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo"
)

func main(){
	app := fiber.New()

	app.Use(cors.New(cors.Config{
        AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
        AllowOrigins:     "*",
        AllowCredentials: true,
        AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
    }))

	// run database
	configs.ConnectDB()

	// collections
	var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")

	api := app.Group("/api/v1")

	// routes
	routes.UserRoute(api, userCollection)

	app.Get("/",func(c *fiber.Ctx) error {
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
		return c.SendString(`<html>
		<body>
			<a href="/api/v1/auth/users/login">Login Google</a>
		</body>
		</html>`)
	})


	app.Listen(":5000")
}