package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
	"github.com/tamhor/lestGo/database"
	"github.com/tamhor/lestGo/router"
)

func main() {
	migrate := flag.Bool("migrate", false, "Databases migration")
	flag.Parse()
	if *migrate {
		database.Migration()
		return
	}

	app := fiber.New()
	database.ConnectDB()
	router.SetupRoutes(app)
	app.Listen(":3000")
}
