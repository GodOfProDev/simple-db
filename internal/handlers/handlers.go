package handlers

import "github.com/godofprodev/simple-db/internal/storage"

type Handlers struct {
	store storage.Storage
}

func New(store storage.Storage) *Handlers {
	return &Handlers{
		store: store,
	}
}
