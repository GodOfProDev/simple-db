package router

import (
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

	app := fiber.New(fiber.Config{
		ErrorHandler: customErrorHandler,
	})

	return &Router{
		app:   app,
		store: store,
	}
}

func (r *Router) RegisterMiddlewares() {
	r.app.Use(cors.New())
	r.app.Use(logger.New(logger.Config{
		Format:     "${cyan}[${time}] ${red}${ip} ${white}${pid} ${red}${status} ${blue}[${method}] ${white}${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "UTC",
	}))
}

func (r *Router) RegisterHandlers() {
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

func (r *Router) Listen(s *config.ServerConfig) error {
	return r.app.Listen(fmt.Sprintf("%v:%v", s.Host, s.Port))
}

func customErrorHandler(c *fiber.Ctx, err error) error {
	switch e := err.(type) {
	case internal.APIError:
		return c.Status(e.Status).JSON(e)
	case internal.APISuccessData:
		return c.Status(e.Status).JSON(e.Data)
	case internal.APISuccessResponse:
		return c.Status(e.Status).JSON(e)
	default:
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{"message": "internal server error"})
	}
}
