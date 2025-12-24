package controller

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Login(c *fiber.Ctx) error {

	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	c.BodyParser(&body)

	if body.Username != "admin" || body.Password != "1234" {
		return c.Status(401).JSON(fiber.Map{"error": "invalid credentials"})
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = body.Username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	signed, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	return c.JSON(fiber.Map{
		"token": signed,
	})
}
