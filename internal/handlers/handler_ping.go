package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h Handlers) HandlePing(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).Send(StringToJson("successfully deleted account id " + uuid.New().String()))
}
