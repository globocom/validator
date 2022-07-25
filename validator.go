package main

import (
	"github.com/go-playground/validator/v10"
	"gitlab.globoi.com/antifraude/validator/validations"
)

func NewValidator() *validator.Validate {
	validate := validator.New()

	// Register all custom validation we create
	validate.RegisterValidation("cpf", validations.CPFValidation)

	return validate
}
