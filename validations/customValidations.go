package validations

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

func CPFValidation(field validator.FieldLevel) bool {
	cpf := field.Field().Interface().(string)

	sum := 0
	aux := 0

	noDots := strings.Replace(cpf, ".", "", -1)
	trimCPF := strings.Replace(noDots, "-", "", -1)

	if trimCPF == "11111111111" || trimCPF == "22222222222" {
		return false
	}

	match, err := regexp.MatchString(`([0-9]{2}[\.]?[0-9]{3}[\.]?[0-9]{3}[\/]?[0-9]{4}[-]?[0-9]{2})|([0-9]{3}[\.]?[0-9]{3}[\.]?[0-9]{3}[-]?[0-9]{2})`, trimCPF)
	if err != nil {
		return false
	}

	if !match {
		return false
	}

	numbers := []rune(trimCPF)
	for i := 1; i <= 9; i++ {
		aux, err = strconv.Atoi(string(numbers[i-1]))
		if err != nil {
			return false
		}
		sum += aux * (11 - i)
	}

	mod := (sum * 10) % 11
	if mod == 10 || mod == 11 {
		mod = 0
	}

	aux, err = strconv.Atoi(string(numbers[9]))
	if err != nil {
		return false
	}

	return aux == mod
}
