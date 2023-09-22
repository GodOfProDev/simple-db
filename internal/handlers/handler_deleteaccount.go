package handlers

import "github.com/gofiber/fiber/v2"

func (h Handlers) HandleDeleteAccount(c *fiber.Ctx) error {
	return c.SendString("delete account")
}
