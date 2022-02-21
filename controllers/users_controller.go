package controllers

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	GetUsers()
}

type userController struct {
	validate *validator.Validate
	path     fiber.Router
}

func NewUserController(api fiber.Router) *userController {
	uc := userController{
		validate: validator.New(),
		path:     api.Group("/users"),
	}

	uc.setupRoutes()

	return &uc
}

func (uc *userController) setupRoutes() {
	uc.path.Get("/", uc.GetUsers())
}

func (uc *userController) GetUsers() fiber.Handler {

	return func(c *fiber.Ctx) error {
		client := fiber.AcquireClient()
		users := client.Get("http://localhost:9010/user/1")
		var response []byte
		code, body, err := users.Get(response, "http://localhost:9010/user/1")

		fmt.Println(code)
		fmt.Println(body)
		fmt.Println(err)

		fmt.Println(users)
		return c.Send(body)
	}
}
