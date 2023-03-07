package main

import (
	"os"
	"tracy-api/configs"
	"tracy-api/ws"

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
	var policeStationCollection *mongo.Collection = configs.GetCollection(configs.DB, "PoliceStations")
	var reportCollection *mongo.Collection = configs.GetCollection(configs.DB, "reports")
	var lokasiCollection *mongo.Collection = configs.GetCollection(configs.DB, "location")

	api := app.Group("/api/v1")

	// routes
	routes.UserRoute(api, userCollection, policeStationCollection)
	routes.PoliceStationRoute(api, policeStationCollection)
	routes.ReportRoute(api,[]*mongo.Collection{reportCollection, userCollection, policeStationCollection})
	routes.LokasiRoute(api,lokasiCollection)

	app.Get("/",func(c *fiber.Ctx) error {
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
		return c.SendString(`<html>
		<body>
			<a href="/api/v1/auth/users/login">Login Google</a>
		</body>
		</html>`)
	})

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)

	go hub.Run()

	app.Post("/ws/createRoom", wsHandler.CreateRoom)
	app.Get("/ws/joinRoom/:roomId", ws.JoinRoom(hub))

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	app.Listen(port)
}