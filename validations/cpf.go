package validations

import (
	"github.com/go-playground/validator/v10"
	"github.com/paemuri/brdoc"
)

func CPFValidation(field validator.FieldLevel) bool {
	cpf := field.Field().Interface().(string)
	return brdoc.IsCPF(cpf)
}
