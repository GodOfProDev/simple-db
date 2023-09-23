package handlers

import (
	"github.com/godofprodev/simple-db/internal"
	"github.com/godofprodev/simple-db/internal/models"
	"github.com/gofiber/fiber/v2"
	"time"
)

func (h Handlers) HandleUpdateUser(c *fiber.Ctx) error {
	uuid, err := getId(c)
	if err != nil {
		return internal.ErrInvalidUUID()
	}

	params := new(models.UpdateUserParams)

	if err := c.BodyParser(params); err != nil {
		return internal.ErrParsingParams()
	}

	if params.Name == "" {
		return internal.ErrNameRequired()
	}

	user, err := h.store.GetUserById(uuid)
	if err != nil {
		return internal.ErrGettingUser()
	}

	user.Name = params.Name
	user.UpdatedAt = time.Now().UTC()

	err = h.store.UpdateUser(user)
	if err != nil {
		return internal.ErrUpdatingUser()
	}

	return internal.SuccessUpdateUser(user)
}
