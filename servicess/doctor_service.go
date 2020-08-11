package servicess

import (
	"docApp/dtos"
	"docApp/models"
	"docApp/repositories"
	"fmt"
	"log"
)

var repoDoctor repositories.DoctorRepo
type DoctorService struct {}

func (*DoctorService) FindById(id string) dtos.Response {
	operationResult := repoDoctor.FindById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: "Can't find id"}
	}

	var data = operationResult
	return dtos.Response{Success: true, Data: data}
}

func (*DoctorService) Create(doctor *models.Doctor) dtos.Response {
	operationResult := repoDoctor.Create(doctor)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.Doctor)

	log.Printf("data at file service: %#v", data)
	return dtos.Response{Success: true, Data: data}
}

func (*DoctorService) FindAll() dtos.Response {
	operationResult := repoDoctor.FindAll()

	if operationResult.Error != nil {
		fmt.Println("file service doctor error")
		log.Println("File service oiiii")
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}


	log.Println("Success get data")
	var data = operationResult
	return dtos.Response{Success: true, Data: data}
}

func (*DoctorService) Update(doctor *models.Doctor) dtos.Response {
	operationResult := repoDoctor.Update(doctor)

	if operationResult.Error != nil {
		return dtos.Response{Success: true, Message: "failed update"}
	}

	var data = operationResult.Result.(*models.Doctor)
	return dtos.Response{Success: true, Data: data}

}

func (*DoctorService) Delete(id string) dtos.Response {
	err := repoDoctor.Delete(id)
	if err.Error != nil{
		return dtos.Response{Success: false, Message: "Failed delete"}
	}

	log.Print("Success Delete")
	return dtos.Response{Success: true, Message: "Success delete"}
}
