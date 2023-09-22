package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h Handlers) HandleGetUser(c *fiber.Ctx) error {
	id := c.Params("id")

	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	user, err := h.store.GetUserById(uid)
	if err != nil {
		return err
	}

	return c.SendString("get account for user with name of " + user.Name)
}
