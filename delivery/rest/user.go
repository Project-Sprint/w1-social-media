package rest

import "github.com/gofiber/fiber/v2"

func UserRegisterHandler(userInterface UserInterface) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("Hello, user register!")
	}
}
