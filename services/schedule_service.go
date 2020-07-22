package service

import (
	"docApp/dtos"
	"docApp/models"
	"docApp/repositories"
	"fmt"
	"log"
)

var repoSchedule repositories.ScheduleRepo
type ScheduleService struct {}

func (*ScheduleService) FindById(id string) dtos.Response {
	operationResult := repoSchedule.FindById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: "Can't find id"}
	}

	var data = operationResult
	return dtos.Response{Success: true, Data: data}
}

func (*ScheduleService) Create(schedule *models.Schedule) dtos.Response {
	operationResult := repoSchedule.Create(schedule)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.Schedule)

	log.Printf("data at file service: %#v", data)
	return dtos.Response{Success: true, Data: data}
}

func (*ScheduleService) FindAll() dtos.Response {
	operationResult := repoSchedule.FindAll()

	if operationResult.Error != nil {
		fmt.Println("file service schedule error")
		log.Println("File service oiiii")
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}


	log.Println("Success get data")
	var data = operationResult.Result
	return dtos.Response{Success: true, Data: data}
}

func (*ScheduleService) Update(schedule *models.Schedule) dtos.Response {
	operationResult := repoSchedule.Update(schedule)

	if operationResult.Error != nil {
		return dtos.Response{Success: true, Message: "failed update"}
	}

	var data = operationResult.Result.(*models.Schedule)
	return dtos.Response{Success: true, Data: data}

}

func (*ScheduleService) Delete(id string) dtos.Response {
	err := repoSchedule.Delete(id)
	if err.Error != nil{
		return dtos.Response{Success: false, Message: "Failed delete"}
	}

	log.Print("Success Delete")
	return dtos.Response{Success: true, Message: "Success delete"}
}
