package helpers

import (
	"docApp/dtos"
	"docApp/langs"

	"gopkg.in/go-playground/validator.v8"
)

func GenerateValidationResponse(err error) (response dtos.ValidationResponse) {
	response.Success = false

	var validations []dtos.ValidationResponse
	validationErrors := err.(validator.ValidationErrors)

	for _, value := range validationErrors {
		field, rule := value.Field, value.Tag
		validation := dtos.Validation{Field: field, Message: langs.GenerateValidationMessage(field, rule)}
		validations = append(validations, validation)
	}
	response.Validations = validations

	return response
}
