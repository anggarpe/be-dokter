package service

import (
	"docApp/models"
	"fmt"
	"log"

	"docApp/dtos"
	"docApp/repositories"
)

func FindById(id string) dtos.Response {
	var repo repositories.UserRepo

	operationResult := repo.FindById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: "Can't find id"}
	}

	var data = operationResult
	return dtos.Response{Success: true, Data: data}
}

func CreateUser(user *models.User) dtos.Response {
	var repo repositories.UserRepo

	//uuidResult, err := uuid.NewRandom()
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//user.ID = uuidResult.String()

	operationResult := repo.Create(user)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.User)

	log.Printf("data at file service: %#v", data)
	return dtos.Response{Success: true, Data: data}
}

func GetAllUser() dtos.Response {
	repo := new(repositories.UserRepo)
	operationResult := repo.FindAll()

	if operationResult.Error != nil {
		fmt.Println("file service user error")
		log.Println("File service oiiii")
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
		}


	log.Println("Success get data")
	var data = operationResult
	return dtos.Response{Success: true, Data: data}
}

func UpdateUser(user *models.User) dtos.Response {
	var repo repositories.UserRepo

	operationResult := repo.Update(user)

	if operationResult.Error != nil {
		return dtos.Response{Success: true, Message: "failed update"}
	}

	var data = operationResult.Result.(*models.User)
	return dtos.Response{Success: true, Data: data}

}

func DeleteUser(id string) dtos.Response {
	var repo repositories.UserRepo
	err := repo.Delete(id)
	if err.Error != nil{
		return dtos.Response{Success: false, Message: "Failed delete"}
	}

	return dtos.Response{Success: true, Message: "Success delete"}
}


