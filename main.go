package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"go-template/controllers"
	"go-template/utils"
)

func main() {
	utils.CreateDbConnection()
	app := fiber.New()
	app.Use(cors.New())

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	controllers.Users(app)
	controllers.Auth(app)
	controllers.Files(app)

	app.Get("/api", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
