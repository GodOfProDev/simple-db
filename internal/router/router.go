package router

import (
	"fmt"
	"github.com/godofprodev/simple-db/internal/config"
	"github.com/godofprodev/simple-db/internal/handlers"
	"github.com/godofprodev/simple-db/internal/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
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

	v1.Get("/metrics", monitor.New())

	v1.Get("/ping", handler.HandlePing)
	v1.Get("/users/:id", handler.HandleGetUser)
	v1.Get("/users", handler.HandleGetUsers)

	v1.Delete("/users/:id", handler.HandleDeleteUser)

	v1.Post("/users", handler.HandleCreateUser)
}

func (r Router) Listen(s *config.ServerConfig) error {
	return r.app.Listen(fmt.Sprintf(":%v", s.Port))

}
