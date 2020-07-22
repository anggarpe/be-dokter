package service

import (
	"docApp/models"
	"fmt"
	"log"

	"docApp/dtos"
	"docApp/repositories"
)
var repoReservation repositories.ReservationRepo

type ReservationService struct {
}


func (*ReservationService) FindById(id string) dtos.Response {
	operationResult := repoReservation.FindById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: "Can't find id"}
	}

	var data = operationResult
	return dtos.Response{Success: true, Data: data}
}

func (*ReservationService) Create(reservation *models.Reservation) dtos.Response {

	//uuidResult, err := uuid.NewRandom()
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//reservation.ID = uuidResult

	operationResult := repoReservation.Create(reservation)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.Reservation)

	log.Printf("data at file service: %#v", data)
	return dtos.Response{Success: true, Data: data}
}

func (*ReservationService) FindAll() dtos.Response {
	operationResult := repoReservation.FindAll()

	if operationResult.Error != nil {
		fmt.Println("file service reservation error")
		log.Println("File service oiiii")
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}


	log.Println("Success get data")
	var data = operationResult
	return dtos.Response{Success: true, Data: data}
}

func (*ReservationService) Update(reservation *models.Reservation) dtos.Response {
	operationResult := repoReservation.Update(reservation)

	if operationResult.Error != nil {
		return dtos.Response{Success: true, Message: "failed update"}
	}

	var data = operationResult.Result.(*models.Reservation)
	return dtos.Response{Success: true, Data: data}

}

func (*ReservationService) Delete(id string) dtos.Response {
	err := repoReservation.Delete(id)
	if err.Error != nil{
		return dtos.Response{Success: false, Message: "Failed delete"}
	}

	log.Print("Success Delete")
	return dtos.Response{Success: true, Message: "Success delete"}
}


