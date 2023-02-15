package routes

import (
	"context"
	"net/http"
	"os"
	"time"
	"tracy-api/configs"
	"tracy-api/controllers"
	"tracy-api/middlewares"
	"tracy-api/repository"
	"tracy-api/services"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func UserRoute(api fiber.Router, userCollection *mongo.Collection, policeStationCollection *mongo.Collection) {

	userRepository := repository.NewUserRepository(userCollection)
	policeStationRepository := repository.NewPoliceStationRepository(policeStationCollection)
	userService := services.NewUserService(userRepository, policeStationRepository)
	userHandler := controllers.NewUserHandler(userService)

	authUser := api.Group("/auth/users")

	authUser.Get("/login",func(c *fiber.Ctx) error {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		defer cancel()
		url := configs.GoogleOAuthConfig().AuthCodeURL(os.Getenv("oAuth_String"))
		
		c.Redirect(url, http.StatusTemporaryRedirect)
		return nil
	})

	authUser.Get("/callback",userHandler.Callback)

	authUser.Get("/profile",middlewares.Auth,userHandler.GetProfile)
	authUser.Put("/profile",middlewares.Auth,userHandler.UpdateProfile)

}