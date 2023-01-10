package validations

import (
	"github.com/go-playground/validator/v10"
	"github.com/paemuri/brdoc"
)

func DocNumberValidation(field validator.FieldLevel) bool {
	docNumber := field.Field().Interface().(string)
	return brdoc.IsCPF(docNumber) || brdoc.IsCNPJ(docNumber)
}
