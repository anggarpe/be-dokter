package servicess

import (
	"docApp/dtos"
	"docApp/models"
	"docApp/repositories"
	"fmt"
	"log"
)

var repoCategory repositories.CategoryRepo
type CategoryService struct {}

func (*CategoryService) FindById(id string) dtos.Response {
	operationResult := repoCategory.FindById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: "Can't find id"}
	}

	var data = operationResult
	return dtos.Response{Success: true, Data: data}
}

func (*CategoryService) Create(category *models.Category) dtos.Response {
	operationResult := repoCategory.Create(category)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.Category)

	log.Printf("data at file service: %#v", data)
	return dtos.Response{Success: true, Data: data}
}

func (*CategoryService) FindAll() dtos.Response {
	operationResult := repoCategory.FindAll()

	if operationResult.Error != nil {
		fmt.Println("file service category error")
		log.Println("File service oiiii")
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}


	log.Println("Success get data")
	var data = operationResult
	return dtos.Response{Success: true, Data: data}
}

func (*CategoryService) Update(category *models.Category) dtos.Response {
	operationResult := repoCategory.Update(category)

	if operationResult.Error != nil {
		return dtos.Response{Success: true, Message: "failed update"}
	}

	var data = operationResult.Result.(*models.Category)
	return dtos.Response{Success: true, Data: data}

}

func (*CategoryService) Delete(id string) dtos.Response {
	err := repoCategory.Delete(id)
	if err.Error != nil{
		return dtos.Response{Success: false, Message: "Failed delete"}
	}

	log.Print("Success Delete")
	return dtos.Response{Success: true, Message: "Success delete"}
}
