package storage

import (
	"github.com/godofprodev/simple-db/internal/models"
	"github.com/google/uuid"
)

type Storage interface {
	CreateUser(user *models.User) error
	DeleteUser(uuid uuid.UUID) error
	UpdateUser(user *models.User) error
	GetUsers() ([]*models.User, error)
	GetUserById(uuid uuid.UUID) (*models.User, error)
}
