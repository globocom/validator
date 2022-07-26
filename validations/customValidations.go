package validations

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

func CPFValidation(field validator.FieldLevel) bool {
	cpf := field.Field().Interface().(string)

	soma := 0
	aux := 0

	noDots := strings.Replace(cpf, ".", "", -1)
	trimCPF := strings.Replace(noDots, "-", "", -1)

	fmt.Println(trimCPF)

	if trimCPF == "11111111111" || trimCPF == "22222222222" {
		fmt.Println("1")
		return false
	}

	match, _ := regexp.MatchString(`([0-9]{2}[\.]?[0-9]{3}[\.]?[0-9]{3}[\/]?[0-9]{4}[-]?[0-9]{2})|([0-9]{3}[\.]?[0-9]{3}[\.]?[0-9]{3}[-]?[0-9]{2})`, trimCPF)
	if match != true {
		return false
	}

	numbers := []rune(trimCPF)
	for i := 1; i <= 9; i++ {
		aux, _ = strconv.Atoi(string(numbers[i-1]))
		soma += aux * (11 - i)
	}

	resto := (soma * 10) % 11
	if resto == 10 || resto == 11 {
		resto = 0
	}

	aux, _ = strconv.Atoi(string(numbers[9]))
	if resto != aux {
		return false
	}

	return true
}
