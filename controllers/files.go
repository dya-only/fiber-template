package controllers

import "github.com/gofiber/fiber/v2"

func GetAvatar(c *fiber.Ctx) error {
	file := c.Params("file")
	return c.SendFile("public/uploads/avatars/" + file)
}
