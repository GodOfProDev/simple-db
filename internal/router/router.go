package router

import (
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

	r.app.Get("/ping", handler.HandlePing)
}

func (r Router) Listen() error {
	return r.app.Listen(":8080")
}
