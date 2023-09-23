package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func (h Handlers) HandleDeleteUser(c *fiber.Ctx) error {
	uuid, err := getId(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).Send(ErrToJson(ErrInvalidUUID))
	}

	err = h.store.DeleteUser(uuid)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Send(ErrToJson(ErrDeletingUser))
	}

	return c.Status(fiber.StatusOK).Send(StringToJson("successfully deleted account id " + uuid.String()))
}
