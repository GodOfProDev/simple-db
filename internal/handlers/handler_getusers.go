package handlers

import (
	"github.com/godofprodev/simple-db/internal"
	"github.com/gofiber/fiber/v2"
)

func (h Handlers) HandleGetUsers(c *fiber.Ctx) error {
	users, err := h.store.GetUsers()
	if err != nil {
		return internal.ErrGettingUsers()
	}

	return internal.SuccessGetUsers(users)
}
