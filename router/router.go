package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	user "github.com/tamhor/lestGo/app/user"
)

func SetupRoutes(app *fiber.App) {
	/**
	Cors Configuration
	*/
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "",
		AllowCredentials: true,
	}))

	api := app.Group("api")
	userApi := api.Group("user")
	userApi.Get("/all", user.GetUsers)
	userApi.Post("/get", user.GetUser)
}
