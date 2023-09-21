package router

import (
	"fmt"
	"github.com/godofprodev/simple-db/internal/config"
	"github.com/godofprodev/simple-db/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	app *fiber.App
}

func New() *Router {
	return &Router{
		app: fiber.New(),
	}
}

func (r Router) AddHandlers() {
	handler := handlers.New()

	v1 := r.app.Group("/v1")

	v1.Get("/ping", handler.HandlePing)
}

func (r Router) Listen(s *config.ServerConfig) error {
	return r.app.Listen(fmt.Sprintf(":%v", s.Port))

}
