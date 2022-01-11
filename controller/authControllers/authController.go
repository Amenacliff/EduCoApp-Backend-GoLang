package authController

import (
	"edu_app_backend/dto"

	"github.com/gofiber/fiber/v2"
)

func LoginUser(context *fiber.Ctx) error {
	context.Accepts("application/json")
	loginRequest := new(dto.LoginRequest)
	err := context.BodyParser(loginRequest)
	if err != nil {
		return err
	}

	return context.JSON(fiber.Map{
		"success": true,
	})
}
