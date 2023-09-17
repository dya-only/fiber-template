package routes

import "github.com/gofiber/fiber/v2"

func Users(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/user", func(c *fiber.Ctx) error {
		return CreateUser(c)
	})

	api.Get("/user", func(c *fiber.Ctx) error {
		return GetAllUser(c)
	})

	api.Get("/user/:id", func(c *fiber.Ctx) error {
		return GetUserById(c)
	})
}
