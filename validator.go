package validator

import (
	"github.com/globocom/validator/validations"
	"github.com/go-playground/validator/v10"
)

func New() *validator.Validate {
	validate := validator.New()

	// Register all custom validation we create
	validate.RegisterValidation("cpf", validations.CPFValidation)

	return validate
}
