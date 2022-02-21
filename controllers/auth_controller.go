package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	validate *validator.Validate
	path     fiber.Router
}

func NewAuthController(api fiber.Router) *AuthController {
	ac := AuthController{
		validate: validator.New(),
		path:     api.Group("/auth"),
	}
	ac.setupRoutes()
	return &ac
}

func (ac *AuthController) setupRoutes() {
	ac.path.Post("/login", ac.login())
}

func (ac *AuthController) login() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("Not Implemented")
	}
}
