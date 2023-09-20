package controllers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	_user "go-template/ent/user"
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
	exists, _ := utils.DbConn.User.Query().Where(_user.NameEQ(body.Name)).Exist(ctx)

	if exists {
		return c.JSON(fiber.Map{
			"success": false,
			"message": "Name already exists",
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
		return c.JSON(fiber.Map{
			"success": false,
			"message": "Failed to find users",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"body":    users,
	})
}

func GetUserById(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	user, err := utils.DbConn.User.Get(ctx, id)

	if user == nil {
		return c.JSON(fiber.Map{
			"success": false,
			"message": "Not found user",
		})
	}
	if err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"message": "Failed to find user by id",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"body":    user,
	})
}

func UpdateAvatar(c *fiber.Ctx) error {
	file, _ := c.FormFile("avatar")
	var body models.Update

	if err := c.BodyParser(&body); err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
		})
	}

	err := c.SaveFile(file, "public/uploads/avatars/"+file.Filename)

	if err != nil || file == nil {
		return c.JSON(fiber.Map{
			"success": false,
			"message": "Failed to save file",
		})
	}

	user, _ := utils.DbConn.User.UpdateOneID(int(body.Id)).SetAvatar(file.Filename).Save(ctx)

	return c.JSON(fiber.Map{
		"success": true,
		"body":    user,
	})
}
