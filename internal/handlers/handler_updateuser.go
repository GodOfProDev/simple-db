package handlers

import (
	"github.com/godofprodev/simple-db/internal/models"
	"github.com/gofiber/fiber/v2"
	"time"
)

func (h Handlers) HandleUpdateUser(c *fiber.Ctx) error {
	uuid, err := getId(c)
	if err != nil {
		return err
	}

	params := new(models.UpdateUserParams)

	if err := c.BodyParser(params); err != nil {
		return err
	}

	if params.Name == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Name is required.")
	}

	user, err := h.store.GetUserById(uuid)
	if err != nil {
		return err
	}

	user.Name = params.Name
	user.UpdatedAt = time.Now().UTC()

	err = h.store.UpdateUser(user)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(user)
}
