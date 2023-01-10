package validations

import (
	"github.com/go-playground/validator/v10"
	"github.com/paemuri/brdoc"
)

func CNPJValidation(field validator.FieldLevel) bool {
	cnpj := field.Field().Interface().(string)
	return brdoc.IsCNPJ(cnpj)
}
