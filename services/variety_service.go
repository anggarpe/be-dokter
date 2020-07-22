package service

import (
	"docApp/dtos"
	"docApp/models"
	"docApp/repositories"
	"fmt"
	"log"
)

var repoVariety repositories.VarietyRepo
type VarietyService struct {}

func (*VarietyService) FindById(id string) dtos.Response {
	operationResult := repoVariety.FindById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: "Can't find id"}
	}

	var data = operationResult
	return dtos.Response{Success: true, Data: data}
}

func (*VarietyService) Create(variety *models.Variety) dtos.Response {
	operationResult := repoVariety.Create(variety)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.Variety)

	log.Printf("data at file service: %#v", data)
	return dtos.Response{Success: true, Data: data}
}

func (*VarietyService) FindAll() dtos.Response {
	operationResult := repoVariety.FindAll()

	if operationResult.Error != nil {
		fmt.Println("file service variety error")
		log.Println("File service oiiii")
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}


	log.Println("Success get data")
	var data = operationResult
	return dtos.Response{Success: true, Data: data}
}

func (*VarietyService) Update(variety *models.Variety) dtos.Response {
	operationResult := repoVariety.Update(variety)

	if operationResult.Error != nil {
		return dtos.Response{Success: true, Message: "failed update"}
	}

	var data = operationResult.Result.(*models.Variety)
	return dtos.Response{Success: true, Data: data}

}

func (*VarietyService) Delete(id string) dtos.Response {
	err := repoVariety.Delete(id)
	if err.Error != nil{
		return dtos.Response{Success: false, Message: "Failed delete"}
	}

	log.Print("Success Delete")
	return dtos.Response{Success: true, Message: "Success delete"}
}
