package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h Handlers) HandleDeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	err = h.store.DeleteUser(uid)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).SendString("Successfully deleted the account")
}
