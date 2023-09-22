package handlers

import (
	"github.com/godofprodev/simple-db/internal/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Handlers struct {
	store storage.Storage
}

func New(store storage.Storage) *Handlers {
	return &Handlers{
		store: store,
	}
}

func getId(c *fiber.Ctx) (uuid.UUID, error) {
	id := c.Params("id")

	uid, err := uuid.Parse(id)
	if err != nil {
		return uuid.UUID{}, err
	}

	return uid, nil
}
