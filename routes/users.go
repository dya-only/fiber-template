package routes

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go-template/models"
	"go-template/utils"
	"time"
)

var ctx = context.Background()

func CreateUser(c *fiber.Ctx) error {
	var body models.User

	if err := c.BodyParser(&body); err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
		})
	}

	user, err := utils.DbConn.User.
		Create().
		SetName(body.Name).
		SetPassword(body.Password).
		SetCreateAt(time.Now()).
		SetUpdateAt(time.Now()).
		Save(ctx)

	if err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"message": "Failed to create user",
			"error":   err,
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"body":    user,
	})
}

func GetAllUser(c *fiber.Ctx) error {
	users, err := utils.DbConn.User.Query().All(ctx)

	if users == nil {
		return c.JSON(fiber.Map{
			"success": false,
			"message": "Not found users",
		})
	}
	if err != nil {
		return c.JSON(err)
	}
	return c.JSON(users)
}

func Users(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/user", func(c *fiber.Ctx) error {
		return CreateUser(c)
	})

	api.Get("/user", func(c *fiber.Ctx) error {
		return GetAllUser(c)
	})
}
