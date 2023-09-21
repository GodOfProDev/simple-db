package main

import (
	"github.com/godofprodev/simple-db/internal/router"
	"log"
)

func main() {
	r := router.New()
	r.AddHandlers()

	err := r.Listen()
	if err != nil {
		log.Fatal("There was an issue listening to port 8080: ", err)
	}
}
