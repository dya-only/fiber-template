package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-template/middlewares"
)

func Users(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/user", func(c *fiber.Ctx) error {
		return CreateUser(c)
	})

	api.Get("/user", middlewares.JwtGuard, func(c *fiber.Ctx) error {
		return GetAllUser(c)
	})

	api.Get("/user/:id", middlewares.JwtGuard, func(c *fiber.Ctx) error {
		return GetUserById(c)
	})

	api.Patch("/user/avatar", middlewares.JwtGuard, func(c *fiber.Ctx) error {
		return UpdateAvatar(c)
	})
}

func Auth(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/auth/by-pass", func(c *fiber.Ctx) error {
		return LoginByPassword(c)
	})
}

func Files(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/uploads/avatar/:file", middlewares.JwtGuard, func(c *fiber.Ctx) error {
		return GetAvatar(c)
	})
}
