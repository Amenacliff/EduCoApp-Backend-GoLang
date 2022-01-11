package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"edu_app_backend/controller"
)

func SetUpRoutes(router fiber.Router) {
	fmt.Println("Hello")
	// context := fiber.Ctx
	userRoute := router.Group("/user")

	userRoute.Post("/create", controller.CreateUser)
}
