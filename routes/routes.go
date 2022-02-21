package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pebeasley/gateway/controllers"
)

func SetupRotues(api fiber.Router) {
	controllers.NewAuthController(api)
	controllers.NewUserController(api)
}
