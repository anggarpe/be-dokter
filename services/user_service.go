package service

import (
	"log"

	"docApp/dtos"
	"docApp/repositories"

	"github.com/google/uuid"
)

func CreateUser(user *models.User, repo repositories.UserRepo) dtos.Response {
	uuidResult, err := uuid.NewRandom()

	if err != nil {
		log.Fatalln(err)
	}

	user.ID = uuidResult

	operationResult := repo.Save(user)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.User)
	return dtos.Response{Success: true, Data: data}
}
