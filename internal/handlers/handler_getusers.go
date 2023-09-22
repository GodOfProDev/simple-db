package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func (h Handlers) HandleGetUsers(c *fiber.Ctx) error {
	users, err := h.store.GetUsers()
	if err != nil {
		return nil
	}

	return c.Status(fiber.StatusOK).JSON(users)
}
