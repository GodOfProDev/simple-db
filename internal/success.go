package internal

import (
	"encoding/json"
	"github.com/godofprodev/simple-db/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type APISuccessData struct {
	Status int         `json:"-"`
	Data   interface{} `json:"-"`
}

type APISuccessResponse struct {
	Status int    `json:"-"`
	Msg    string `json:"message"`
}

func (e APISuccessData) Error() string {
	jsonBytes, _ := json.Marshal(e.Data)
	return string(jsonBytes)
}

func (e APISuccessResponse) Error() string {
	return e.Msg
}

func SuccessDeleteUser(uuid uuid.UUID) APISuccessResponse {
	return APISuccessResponse{
		Status: fiber.StatusOK,
		Msg:    "successfully deleted account id " + uuid.String(),
	}
}

func SuccessCreateUser(user *models.User) APISuccessData {
	return APISuccessData{
		Status: fiber.StatusCreated,
		Data:   user,
	}
}

func SuccessGetUser(user *models.User) APISuccessData {
	return APISuccessData{
		Status: fiber.StatusOK,
		Data:   user,
	}
}

func SuccessGetUsers(users []*models.User) APISuccessData {
	return APISuccessData{
		Status: fiber.StatusOK,
		Data:   users,
	}
}

func SuccessUpdateUser(user *models.User) APISuccessData {
	return APISuccessData{
		Status: fiber.StatusOK,
		Data:   user,
	}
}
