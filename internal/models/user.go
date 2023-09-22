package models

import (
	"github.com/google/uuid"
	"time"
)

type CreateUserParams struct {
	Name string `json:"name"`
}

type User struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewUser(p *CreateUserParams) *User {
	return &User{
		Id:        uuid.New(),
		Name:      p.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
}
