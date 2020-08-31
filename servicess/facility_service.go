package servicess

import (
	"docApp/dtos"
	"docApp/models"
	"docApp/repositories"
	"fmt"
	"log"
)

var repoFacility repositories.FacilityRepo
type FacilityService struct {}

func (*FacilityService) FindById(id string) dtos.Response {
	operationResult := repoFacility.FindById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: "Can't find id"}
	}

	var data = operationResult
	return dtos.Response{Success: true, Data: data}
}

func (*FacilityService) Create(facility *models.Facility) dtos.Response {
	operationResult := repoFacility.Create(facility)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.Facility)

	log.Printf("data at file service: %#v", data)
	return dtos.Response{Success: true, Data: data}
}

func (*FacilityService) FindAll() dtos.Response {
	operationResult := repoFacility.FindAll()

	if operationResult.Error != nil {
		fmt.Println("file service facility error")
		log.Println("File service oiiii")
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}


	log.Println("Success get data")
	var data = operationResult
	return dtos.Response{Success: true, Data: data}
}

func (*FacilityService) Update(facility *models.Facility) dtos.Response {
	operationResult := repoFacility.Update(facility)

	if operationResult.Error != nil {
		return dtos.Response{Success: true, Message: "failed update"}
	}

	var data = operationResult.Result.(*models.Facility)
	return dtos.Response{Success: true, Data: data}

}

func (*FacilityService) Delete(id string) dtos.Response {
	err := repoFacility.Delete(id)
	if err.Error != nil{
		return dtos.Response{Success: false, Message: "Failed delete"}
	}

	log.Print("Success Delete")
	return dtos.Response{Success: true, Message: "Success delete"}
}
