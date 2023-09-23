package internal

import (
	"github.com/gofiber/fiber/v2"
)

type APIError struct {
	Status int    `json:"-"`
	Msg    string `json:"message"`
}

func (e APIError) Error() string {
	return e.Msg
}

func ErrNameRequired() APIError {
	return APIError{
		Status: fiber.StatusBadRequest,
		Msg:    "name is required",
	}
}

func ErrCreatingUser() APIError {
	return APIError{
		Status: fiber.StatusInternalServerError,
		Msg:    "there was an issue creating the user",
	}
}

func ErrDeletingUser() APIError {
	return APIError{
		Status: fiber.StatusInternalServerError,
		Msg:    "there was an issue deleting the user",
	}
}

func ErrGettingUser() APIError {
	return APIError{
		Status: fiber.StatusInternalServerError,
		Msg:    "there was an issue getting the user",
	}
}

func ErrGettingUsers() APIError {
	return APIError{
		Status: fiber.StatusInternalServerError,
		Msg:    "there was an issue getting the users",
	}
}

func ErrUpdatingUser() APIError {
	return APIError{
		Status: fiber.StatusInternalServerError,
		Msg:    "there was an issue updating the users",
	}
}

func ErrInvalidUUID() APIError {
	return APIError{
		Status: fiber.StatusBadRequest,
		Msg:    "invalid uuid",
	}
}

func ErrParsingParams() APIError {
	return APIError{
		Status: fiber.StatusBadRequest,
		Msg:    "there was an issue parsing params",
	}
}
