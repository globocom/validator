package validator

import (
	"log"

	"github.com/globocom/validator/validations"
	"github.com/go-playground/validator/v10"
)

func New() *validator.Validate {
	validate := validator.New()

	if err := validate.RegisterValidation("cpf", validations.CPFValidation); err != nil {
		log.Fatal(err)
	}

	if err := validate.RegisterValidation("cnpj", validations.CNPJValidation); err != nil {
		log.Fatal(err)
	}

	if err := validate.RegisterValidation("docNumber", validations.DocNumberValidation); err != nil {
		log.Fatal(err)
	}

	return validate
}
