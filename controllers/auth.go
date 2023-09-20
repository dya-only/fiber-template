package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	_user "go-template/ent/user"
	"go-template/models"
	"go-template/utils"
	"os"
	"time"
)

func LoginByPassword(c *fiber.Ctx) error {
	var body models.User

	if err := c.BodyParser(&body); err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
		})
	}

	id, _ := utils.DbConn.User.Query().Where(_user.NameEQ(body.Name)).OnlyID(ctx)

	if id == 0 {
		return c.JSON(fiber.Map{
			"success": false,
			"message": "Not found user",
		})
	}

	user, _ := utils.DbConn.User.Get(ctx, id)

	if body.Password != user.Password {
		return c.JSON(fiber.Map{
			"success": false,
			"message": "Failed to login by password",
		})
	}

	claims := jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	_token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := _token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"message": "Failed create token",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"token":   token,
	})
}
