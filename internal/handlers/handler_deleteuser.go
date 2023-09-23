package handlers

import (
	"github.com/godofprodev/simple-db/internal"
	"github.com/gofiber/fiber/v2"
)

func (h Handlers) HandleDeleteUser(c *fiber.Ctx) error {
	uuid, err := getId(c)
	if err != nil {
		return internal.ErrInvalidUUID()
	}

	err = h.store.DeleteUser(uuid)
	if err != nil {
		return internal.ErrDeletingUser()
	}

	return internal.SuccessDeleteUser(uuid)
}
