package main

import (
	"fmt"
	"github.com/godofprodev/simple-db/internal/config"
	"github.com/godofprodev/simple-db/internal/router"
	"github.com/godofprodev/simple-db/internal/storage"
	"log/slog"
)

func main() {
	v, err := initViper()
	if err != nil {
		slog.Error("failed to initialize viper: ", err)
	}

	serverCfg := config.NewServerConfig(v)
	dbCfg := config.NewDBConfig(v)

	db, err := storage.NewPostgresStore(dbCfg)
	if err != nil {
		slog.Error("there was an issue connecting to postgres db: ", err)
	}

	r := router.New(db)
	r.AddHandlers()

	err = r.Listen(serverCfg)
	if err != nil {
		slog.Error(fmt.Sprintf("there was an issue listening to port : %v", v.GetInt("PORT")), err)
	}
}
