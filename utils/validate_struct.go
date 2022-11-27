package utils

import "github.com/go-playground/validator/v10"

type ErrorSchema struct {
	FailedField string `json:"failedField"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
}

func ValidateStruct(body any) []*ErrorSchema {
	var validate = validator.New()
	var errors []*ErrorSchema
	err := validate.Struct(body)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorSchema
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
