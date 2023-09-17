package routes

import "github.com/gofiber/fiber/v2"

func GetAvatar(c *fiber.Ctx) error {
	file := c.Params("fileName")
	return c.SendFile("public/uploads/avatars/" + file)
}
