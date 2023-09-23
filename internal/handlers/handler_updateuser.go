package handlers

import (
	"github.com/godofprodev/simple-db/internal/models"
	"github.com/gofiber/fiber/v2"
	"time"
)

func (h Handlers) HandleUpdateUser(c *fiber.Ctx) error {
	uuid, err := getId(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).Send(ErrToJson(ErrInvalidUUID))
	}

	params := new(models.UpdateUserParams)

	if err := c.BodyParser(params); err != nil {
		return c.Status(fiber.StatusBadRequest).Send(ErrToJson(ErrParsingParams))
	}

	if params.Name == "" {
		return c.Status(fiber.StatusBadRequest).Send(ErrToJson(ErrNameRequired))
	}

	user, err := h.store.GetUserById(uuid)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Send(ErrToJson(ErrGettingUser))
	}

	user.Name = params.Name
	user.UpdatedAt = time.Now().UTC()

	err = h.store.UpdateUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Send(ErrToJson(ErrUpdatingUser))
	}

	return c.Status(fiber.StatusOK).JSON(user)
}
