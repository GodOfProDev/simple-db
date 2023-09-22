package handlers

import "github.com/gofiber/fiber/v2"

func (h Handlers) HandleCreateAccount(c *fiber.Ctx) error {
	return c.SendString("create account")
}
