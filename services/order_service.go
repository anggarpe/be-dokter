package service

import (
	"docApp/models"
	"fmt"
	"log"

	"docApp/dtos"
	"docApp/repositories"
)
var repoOrder repositories.OrderRepo

type OrderService struct {
}


func (*OrderService) FindById(id string) dtos.Response {
	operationResult := repoOrder.FindById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: "Can't find id"}
	}

	var data = operationResult
	return dtos.Response{Success: true, Data: data}
}

func (*OrderService) Create(order *models.Order) dtos.Response {

	//uuidResult, err := uuid.NewRandom()
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//order.ID = uuidResult

	operationResult := repoOrder.Create(order)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.Order)

	log.Printf("data at file service: %#v", data)
	return dtos.Response{Success: true, Data: data}
}

func (*OrderService) FindAll() dtos.Response {
	operationResult := repoOrder.FindAll()

	if operationResult.Error != nil {
		fmt.Println("file service order error")
		log.Println("File service oiiii")
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}


	log.Println("Success get data")
	var data = operationResult
	return dtos.Response{Success: true, Data: data}
}

func (*OrderService) Update(order *models.Order) dtos.Response {
	operationResult := repoOrder.Update(order)

	if operationResult.Error != nil {
		return dtos.Response{Success: true, Message: "failed update"}
	}

	var data = operationResult.Result.(*models.Order)
	return dtos.Response{Success: true, Data: data}

}

func (*OrderService) Delete(id string) dtos.Response {
	err := repoOrder.Delete(id)
	if err.Error != nil{
		return dtos.Response{Success: false, Message: "Failed delete"}
	}

	log.Print("Success Delete")
	return dtos.Response{Success: true, Message: "Success delete"}
}


