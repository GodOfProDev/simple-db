package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func (h Handlers) HandleGetUser(c *fiber.Ctx) error {
	uuid, err := getId(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).Send(ErrToJson(ErrInvalidUUID))
	}

	user, err := h.store.GetUserById(uuid)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Send(ErrToJson(ErrGettingUser))
	}

	return c.Status(fiber.StatusOK).JSON(user)
}
