package routes

import (
	"context"
	"net/http"
	"os"
	"time"
	"tracy-api/configs"
	"tracy-api/controllers"
	"tracy-api/repository"
	"tracy-api/services"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func UserRoute(api fiber.Router, userCollection *mongo.Collection) {

	userRepository := repository.NewUserRepository(userCollection)
	userService := services.NewUserService(userRepository)
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

}