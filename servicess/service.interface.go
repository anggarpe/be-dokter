package servicess

import (
	"docApp/dtos"
)

type Service interface {
	FindById(id string) dtos.Response
	Create(interface{}) dtos.Response
	FindAll() 			dtos.Response
	Update(interface{}) dtos.Response
	Delete(id string) 	dtos.Response
}
