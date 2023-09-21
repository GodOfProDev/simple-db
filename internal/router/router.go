package router

import (
	"fmt"
	"github.com/godofprodev/simple-db/internal/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
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

func (r Router) Listen(v *viper.Viper) error {
	return r.app.Listen(fmt.Sprintf(":%v", v.GetString("PORT")))

}
