package handlers

import (
	"errors"
)

var (
	ErrNameRequired  = errors.New("name is required")
	ErrCreatingUser  = errors.New("there was an issue creating the user")
	ErrDeletingUser  = errors.New("there was an issue deleting the user")
	ErrGettingUser   = errors.New("there was an issue getting the user")
	ErrGettingUsers  = errors.New("there was an issue getting the users")
	ErrUpdatingUser  = errors.New("there was an issue updating the users")
	ErrInvalidUUID   = errors.New("invalid uuid")
	ErrParsingParams = errors.New("there was an issue parsing params")
)
