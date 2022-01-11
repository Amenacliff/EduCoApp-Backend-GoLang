package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	authController "edu_app_backend/controller/authControllers"
	"edu_app_backend/controller/userController"
)

func SetUpRoutes(router fiber.Router) {
	fmt.Println("Hello")
	// context := fiber.Ctx
	userRoute := router.Group("/user")
	authRoute := router.Group("/auth")

	userRoute.Post("/create", userController.CreateUser)
	authRoute.Post("/login", authController.LoginUser)
}
