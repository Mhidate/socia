package main

import (
	"socia/config"
	"socia/internal/http/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	config.ConnectDB()

	routes.AuthRoutes(app)
	routes.PhotoRoutes(app)

	app.Listen(":3000")
}
