package handlers

import (
	"github.com/godofprodev/simple-db/internal"
	"github.com/gofiber/fiber/v2"
)

func (h Handlers) HandlePing(c *fiber.Ctx) error {
	return internal.ErrNameRequired()
}
