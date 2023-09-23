package handlers

import (
	"github.com/godofprodev/simple-db/internal"
	"github.com/gofiber/fiber/v2"
)

func (h Handlers) HandleGetUser(c *fiber.Ctx) error {
	uuid, err := getId(c)
	if err != nil {
		return internal.ErrInvalidUUID()
	}

	user, err := h.store.GetUserById(uuid)
	if err != nil {
		return internal.ErrGettingUser()
	}

	return internal.SuccessGetUser(user)
}
