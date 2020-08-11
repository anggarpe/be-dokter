package servicess

import (
	"docApp/dtos"
	"docApp/models"
	"docApp/repositories"
	"fmt"
	"log"
)

var repoProduct repositories.ProductRepo
type ProductService struct {}

func (*ProductService) FindById(id string) dtos.Response {
	operationResult := repoProduct.FindById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: "Can't find id"}
	}

	var data = operationResult
	return dtos.Response{Success: true, Data: data}
}

func (*ProductService) Create(product *models.Product) dtos.Response {
	operationResult := repoProduct.Create(product)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.Product)

	log.Printf("data at file service: %#v", data)
	return dtos.Response{Success: true, Data: data}
}

func (*ProductService) FindAll() dtos.Response {
	operationResult := repoProduct.FindAll()

	if operationResult.Error != nil {
		fmt.Println("file service product error")
		log.Println("File service oiiii")
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}


	log.Println("Success get data")
	var data = operationResult
	return dtos.Response{Success: true, Data: data}
}

func (*ProductService) Update(product *models.Product) dtos.Response {
	operationResult := repoProduct.Update(product)

	if operationResult.Error != nil {
		return dtos.Response{Success: true, Message: "failed update"}
	}

	var data = operationResult.Result.(*models.Product)
	return dtos.Response{Success: true, Data: data}

}

func (*ProductService) Delete(id string) dtos.Response {
	err := repoProduct.Delete(id)
	if err.Error != nil{
		return dtos.Response{Success: false, Message: "Failed delete"}
	}

	log.Print("Success Delete")
	return dtos.Response{Success: true, Message: "Success delete"}
}

func (*ProductService) FindAllByPrice(min, max float64) dtos.Response {
	operationResult := repoProduct.FindAllByPrice(min, max)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}


	log.Println("Success get data")
	var data = operationResult
	return dtos.Response{Success: true, Data: data}
}

func (*ProductService) Search(name string) dtos.Response {
	operationResult := repoProduct.Search(name)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	log.Println("Success get data")
	var data = operationResult
	return dtos.Response{Success: true, Data: data}
}

func (*ProductService) OrderByPriceDesc() dtos.Response {
	operationResult := repoProduct.OrderByPriceDsc()

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	log.Println("Success get data")
	var data = operationResult
	return dtos.Response{Success: true, Data: data}
}
