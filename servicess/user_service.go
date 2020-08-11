package servicess

import (
	"docApp/models"
	"fmt"
	"log"

	"docApp/dtos"
	"docApp/repositories"
)
var repoUser repositories.UserRepo

type UserService struct {
}


func (*UserService) FindById(id string) dtos.Response {
	operationResult := repoUser.FindById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: "Can't find id"}
	}

	var data = operationResult
	return dtos.Response{Success: true, Data: data}
}

func (*UserService) Create(user *models.User) dtos.Response {

	//uuidResult, err := uuid.NewRandom()
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//user.ID = uuidResult

	operationResult := repoUser.Create(user)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.User)

	log.Printf("data at file service: %#v", data)
	return dtos.Response{Success: true, Data: data}
}

func (*UserService) FindAll() dtos.Response {
	operationResult := repoUser.FindAll()

	if operationResult.Error != nil {
		fmt.Println("file service user error")
		log.Println("File service oiiii")
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
		}


	log.Println("Success get data")
	var data = operationResult
	return dtos.Response{Success: true, Data: data}
}

func (*UserService) Update(user *models.User) dtos.Response {
	operationResult := repoUser.Update(user)

	if operationResult.Error != nil {
		return dtos.Response{Success: true, Message: "failed update"}
	}

	var data = operationResult.Result.(*models.User)
	return dtos.Response{Success: true, Data: data}

}

func (*UserService) Delete(id string) dtos.Response {
	err := repoUser.Delete(id)
	if err.Error != nil{
		return dtos.Response{Success: false, Message: "Failed delete"}
	}

	log.Print("Success Delete")
	return dtos.Response{Success: true, Message: "Success delete"}
}


