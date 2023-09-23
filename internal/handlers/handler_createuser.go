package handlers

import (
	"github.com/godofprodev/simple-db/internal/models"
	"github.com/gofiber/fiber/v2"
)

func (h Handlers) HandleCreateUser(c *fiber.Ctx) error {
	params := new(models.CreateUserParams)

	if err := c.BodyParser(params); err != nil {
		return c.Status(fiber.StatusBadRequest).Send(ErrToJson(ErrParsingParams))
	}

	if params.Name == "" {
		return c.Status(fiber.StatusBadRequest).Send(ErrToJson(ErrNameRequired))
	}

	user := models.NewUser(params)

	if err := h.store.CreateUser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).Send(ErrToJson(ErrCreatingUser))
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}
