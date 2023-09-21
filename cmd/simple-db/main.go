package main

import (
	"fmt"
	"github.com/godofprodev/simple-db/internal/config"
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

	serverCfg := config.NewServerConfig(v)

	err = r.Listen(serverCfg)
	if err != nil {
		slog.Error(fmt.Sprintf("there was an issue listening to port : %v", v.GetInt("PORT")), err)
	}
}
