package servicess

import (
	"docApp/dtos"
	"docApp/models"
	"docApp/repositories"
	"fmt"
	"log"
)


var repoAdmin repositories.AdminRepo
type AdminService struct {}

func (*AdminService) FindById(id string) dtos.Response {
	operationResult := repoAdmin.FindById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: "Can't find id"}
	}

	var data = operationResult
	return dtos.Response{Success: true, Data: data}
}

func (*AdminService) Create(user *models.Admin) dtos.Response {
	operationResult := repoAdmin.Create(user)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.Admin)

	//log.Printf("data at file service: %#v", data)
	return dtos.Response{Success: true, Data: data}
}

func (*AdminService) FindAll() dtos.Response {
	operationResult := repoAdmin.FindAll()

	if operationResult.Error != nil {
		fmt.Println("file service user error")
		log.Println("File service nih")
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}


	log.Println("Success get data")
	var data = operationResult
	return dtos.Response{Success: true, Data: data}
}

func (*AdminService) Update(user *models.Admin) dtos.Response {
	operationResult := repoAdmin.Update(user)

	if operationResult.Error != nil {
		return dtos.Response{Success: true, Message: "failed update"}
	}

	var data = operationResult.Result.(*models.Admin)
	return dtos.Response{Success: true, Data: data}

}

func (*AdminService) Delete(id string) dtos.Response {
	err := repoAdmin.Delete(id)
	if err.Error != nil{
		return dtos.Response{Success: false, Message: "Failed delete"}
	}

	log.Print("Success Delete")
	return dtos.Response{Success: true, Message: "Success delete"}
}
