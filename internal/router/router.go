package router

import (
	"errors"
	"fmt"
	"github.com/godofprodev/simple-db/internal"
	"github.com/godofprodev/simple-db/internal/config"
	"github.com/godofprodev/simple-db/internal/handlers"
	"github.com/godofprodev/simple-db/internal/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

type Router struct {
	app   *fiber.App
	store storage.Storage
}

func New(store storage.Storage) *Router {

	app := fiber.New(fiber.Config{ErrorHandler: func(c *fiber.Ctx, err error) error {
		var apiErr internal.APIError
		if errors.As(err, &apiErr) {
			return c.Status(apiErr.Status).JSON(apiErr)
		}

		var apiSuccessData internal.APISuccessData
		if errors.As(err, &apiSuccessData) {
			return c.Status(apiSuccessData.Status).JSON(apiSuccessData.Data)
		}

		var apiSuccessString internal.APISuccessString
		if errors.As(err, &apiSuccessString) {
			return c.Status(apiSuccessString.Status).JSON(apiSuccessString)
		}

		return c.Status(fiber.StatusInternalServerError).JSON(map[string]any{"message": "internal server error"})
	}})

	return &Router{
		app:   app,
		store: store,
	}
}

func (r Router) RegisterMiddlewares() {
	r.app.Use(cors.New())
	r.app.Use(logger.New(logger.Config{
		Format:     "${cyan}[${time}] ${red}${ip} ${white}${pid} ${red}${status} ${blue}[${method}] ${white}${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "UTC",
	}))
}

func (r Router) RegisterHandlers() {
	handler := handlers.New(r.store)

	v1 := r.app.Group("/v1")

	v1.Get("/monitor", monitor.New())

	v1.Get("/ping", handler.HandlePing)
	v1.Get("/users/:id", handler.HandleGetUser)
	v1.Get("/users", handler.HandleGetUsers)

	v1.Delete("/users/:id", handler.HandleDeleteUser)

	v1.Post("/users", handler.HandleCreateUser)

	v1.Put("/users/:id", handler.HandleUpdateUser)
}

func (r Router) Listen(s *config.ServerConfig) error {
	return r.app.Listen(fmt.Sprintf("%v:%v", s.Host, s.Port))
}
