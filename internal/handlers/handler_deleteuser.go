package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func (h Handlers) HandleDeleteUser(c *fiber.Ctx) error {
	uuid, err := getId(c)
	if err != nil {
		return err
	}

	err = h.store.DeleteUser(uuid)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).SendString("Successfully deleted the account with id of " + uuid.String())
}
