package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/pebeasley/gateway/routes"
)

type App struct {
	*fiber.App
}

func main() {
	fmt.Println("hello from gateway")

	app := App{
		App: fiber.New(),
	}

	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(recover.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello world")
	})

	routes.SetupRotues(app)

	err := app.Listen(":6080")

	if err != nil {
		fmt.Println(err.Error())
	}
}
