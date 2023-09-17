package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go-template/routes"
	"go-template/utils"
)

func main() {
	utils.CreateDbConnection()
	app := fiber.New()
	app.Use(cors.New())

	routes.Users(app)

	app.Get("/api", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
