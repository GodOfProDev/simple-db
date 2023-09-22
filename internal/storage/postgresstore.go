package storage

import (
	"database/sql"
	"github.com/godofprodev/simple-db/internal/config"
	"github.com/godofprodev/simple-db/internal/models"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type PostgresStore struct {
	DB *sql.DB
}

func NewPostgresStore(cfg *config.DBConfig) (*PostgresStore, error) {
	db, err := sql.Open("postgres", cfg.Url)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		DB: db,
	}, nil
}

func (s *PostgresStore) CreateUser(user *models.User) error {
	//TODO implement me
	panic("implement me")
}

func (s *PostgresStore) DeleteUser(uuid uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (s *PostgresStore) UpdateUser(user *models.User) error {
	//TODO implement me
	panic("implement me")
}

func (s *PostgresStore) GetUserById(uuid uuid.UUID) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}
