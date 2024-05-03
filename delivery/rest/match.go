package rest

import "github.com/gofiber/fiber/v2"

func PostMatchHandler(matchInterface MatchService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("Hello, Match Post!")
	}
}

func ApproveMatchHandler(matchInterface MatchService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("Hello, Match Approve!")
	}
}

func RejectMatchHandler(matchInterface MatchService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("Hello, Match Reject!")
	}
}
