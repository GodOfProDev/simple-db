package storage

import (
	"database/sql"
	"errors"
	"github.com/godofprodev/simple-db/internal/config"
	"github.com/godofprodev/simple-db/internal/models"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type PostgresStore struct {
	DB *sql.DB
}

const (
	createUserSQL  = `INSERT INTO users VALUES ($1, $2, $3, $4)`
	deleteUserSQL  = `DELETE FROM users WHERE id = $1`
	getUsersSQL    = `SELECT * FROM users`
	getUserByIdSQL = `SELECT * FROM users WHERE id = $1`
)

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
	_, err := s.DB.Exec(createUserSQL, user.Id, user.Name, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStore) DeleteUser(uuid uuid.UUID) error {
	_, err := s.DB.Exec(deleteUserSQL, uuid)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStore) UpdateUser(user *models.User) error {
	//TODO implement me
	panic("implement me")
}

func (s *PostgresStore) GetUsers() ([]*models.User, error) {
	rows, err := s.DB.Query(getUsersSQL)
	if err != nil {
		return nil, err
	}

	var users []*models.User

	for rows.Next() {
		user, err := scanIntoUsers(rows)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (s *PostgresStore) GetUserById(uuid uuid.UUID) (*models.User, error) {
	rows, err := s.DB.Query(getUserByIdSQL, uuid)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoUsers(rows)
	}

	return nil, errors.New("account not found")
}

func scanIntoUsers(rows *sql.Rows) (*models.User, error) {
	user := new(models.User)

	err := rows.Scan(
		&user.Id,
		&user.Name,
		&user.CreatedAt,
		&user.UpdatedAt)

	return user, err
}
