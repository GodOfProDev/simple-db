package router

import (
	handlers "github.com/godofprodev/simple-db/internal/handlers"
	"github.com/gofiber/fiber/v2"
	"log"
)

type Router struct{}

func New() *Router {
	return &Router{}
}

func (r Router) AddHandlers() {
	app := fiber.New()

	handler := handlers.New()

	app.Get("/ping", handler.HandlePing)

	err := app.Listen(":8080")
	if err != nil {
		log.Fatal("There was an issue listening to port 8080: ", err)
	}
}
