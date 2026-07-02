package utils

import "github.com/go-playground/validator/v10"

var Validate *validator.Validate

func init() {
	Validate = NewValidator()
}


func NewValidator() *validator.Validate{
	return validator.New(validator.WithRequiredStructEnabled())
}
