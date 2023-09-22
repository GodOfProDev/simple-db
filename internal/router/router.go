package router

import (
	"fmt"
	"github.com/godofprodev/simple-db/internal/config"
	"github.com/godofprodev/simple-db/internal/handlers"
	"github.com/godofprodev/simple-db/internal/storage"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	app   *fiber.App
	store storage.Storage
}

func New(store storage.Storage) *Router {
	return &Router{
		app:   fiber.New(),
		store: store,
	}
}

func (r Router) AddHandlers() {
	handler := handlers.New(r.store)

	v1 := r.app.Group("/v1")

	v1.Get("/ping", handler.HandlePing)
	v1.Get("/user/:id", handler.HandleGetAccount)
	v1.Delete("/user/:id", handler.HandleDeleteAccount)
	v1.Post("/user", handler.HandleCreateAccount)
}

func (r Router) Listen(s *config.ServerConfig) error {
	return r.app.Listen(fmt.Sprintf(":%v", s.Port))

}
