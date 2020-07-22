package service

import (
	"docApp/dtos"
	"docApp/models"
	"docApp/repositories"
	"fmt"
	"log"
)

var repoServices repositories.ServiceRepo
type ServicesService struct {}

func (*ServicesService) FindById(id string) dtos.Response {
	operationResult := repoServices.FindById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: "Can't find id"}
	}

	var data = operationResult
	return dtos.Response{Success: true, Data: data}
}

func (*ServicesService) Create(services *models.Services) dtos.Response {
	operationResult := repoServices.Create(services)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.Services)

	log.Printf("data at file service: %#v", data)
	return dtos.Response{Success: true, Data: data}
}

func (*ServicesService) FindAll() dtos.Response {
	operationResult := repoServices.FindAll()

	if operationResult.Error != nil {
		fmt.Println("file service services error")
		log.Println("File service oiiii")
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}


	log.Println("Success get data")
	var data = operationResult
	return dtos.Response{Success: true, Data: data}
}

func (*ServicesService) Update(services *models.Services) dtos.Response {
	operationResult := repoServices.Update(services)

	if operationResult.Error != nil {
		return dtos.Response{Success: true, Message: "failed update"}
	}

	var data = operationResult.Result.(*models.Services)
	return dtos.Response{Success: true, Data: data}

}

func (*ServicesService) Delete(id string) dtos.Response {
	err := repoServices.Delete(id)
	if err.Error != nil{
		return dtos.Response{Success: false, Message: "Failed delete"}
	}

	log.Print("Success Delete")
	return dtos.Response{Success: true, Message: "Success delete"}
}
