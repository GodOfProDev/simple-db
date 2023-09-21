package main

import (
	"github.com/godofprodev/simple-db/internal/router"
	"log/slog"
)

func main() {
	r := router.New()
	r.AddHandlers()

	v, err := initViper()
	if err != nil {
		slog.Error("failed to initialize viper: ", err)
	}

	err = r.Listen(v)
	if err != nil {
		slog.Error("there was an issue listening to port 8080: ", err)
	}
}
