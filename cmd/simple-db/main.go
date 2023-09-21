package main

import (
	"github.com/godofprodev/simple-db/internal/router"
)

func main() {
	r := router.New()

	r.AddHandlers()
}
