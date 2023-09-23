package handlers

import (
	"github.com/godofprodev/simple-db/internal"
	"github.com/godofprodev/simple-db/internal/models"
	"github.com/gofiber/fiber/v2"
)

func (h Handlers) HandleCreateUser(c *fiber.Ctx) error {
	params := new(models.CreateUserParams)

	if err := c.BodyParser(params); err != nil {
		return internal.ErrParsingParams()
	}

	if params.Name == "" {
		return internal.ErrNameRequired()
	}

	user := models.NewUser(params)

	if err := h.store.CreateUser(user); err != nil {
		return internal.ErrCreatingUser()
	}

	return internal.SuccessCreateUser(user)
}
