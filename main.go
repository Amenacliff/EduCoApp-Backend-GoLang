package main

import (
	"edu_app_backend/routes"
	"edu_app_backend/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	service.ConnectDb()
	routes.SetUpRoutes(app)
	app.Listen(":3000")
}
