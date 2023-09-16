package routes

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go-template/utils"
	"log"
)

var ctx = context.Background()

func GetAllUser(c *fiber.Ctx) error {
	users, err := utils.DbConn.User.Query().All(ctx)

	if users == nil {
		return c.JSON("Not found")
	}
	if err != nil {
		return c.JSON("Err")
	}

	log.Println(users)
	return c.JSON(users)
}

func Users(app *fiber.App) {
	app.Get("/api/user", func(c *fiber.Ctx) error {
		return c.JSON(GetAllUser(c))
	})
}
