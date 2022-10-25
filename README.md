# Validator

<img align="right" src="https://raw.githubusercontent.com/go-playground/validator/v9/logo.png">![Project status](https://img.shields.io/badge/version-1.0.0-green.svg)
[![GoDoc](https://godoc.org/github.com/go-playground/validator?status.svg)](https://pkg.go.dev/github.com/go-playground/validator/v10)

Validator é uma biblioteca interna desenvolvida para criar validações customizadas para as estruturas utilizadas nas aplicações em Go.

Para seu completo entendimento consultar a documentação oficial do [Validator](https://github.com/go-playground/validator/blob/master/README.md).

# Modos de uso

O Validator retorna por padrão uma instância do Validator original, portanto torna acessivel todos os métodos existentes na mesma.

## Chamada Direta

Desta forma você pode chamar a validação a qualquer momento do código, para garantir os requisições de uma estrutura específica.

```golang
type User struct {
	FirstName      string     `validate:"required"`
	LastName       string     `validate:"required"`
	Age            uint8      `validate:"gte=0,lte=130"`
	Email          string     `validate:"required,email"`
	FavouriteColor string     `validate:"iscolor"`
}

func main() {
    validate = validator.New()

    user := &User{
		FirstName:      "Badger",
		LastName:       "Smith",
		Age:            135,
		Email:          "Badger.Smith@gmail.com",
		FavouriteColor: "#000-",
	}

	if err := validate.Struct(user); err != nil {
        fmt.Println("Error when validate - ", err)
        panic(err)
    }
}
```

## Encapsulada no Echo Framework

A framework Echo permite que definamos um validator padrão para todas as requisições do mesmo, desta maneira, própria instância do Echo terá o validator internamente.

```golang
type User struct {
	FirstName      string     `validate:"required"`
	LastName       string     `validate:"required"`
	Age            uint8      `validate:"gte=0,lte=130"`
	Email          string     `validate:"required,email"`
	FavouriteColor string     `validate:"iscolor"`
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
    e := echo.New()
    e.Validator = &CustomValidator{validator: validator.New()}

    user := &User{
		FirstName:      "Badger",
		LastName:       "Smith",
		Age:            135,
		Email:          "Badger.Smith@gmail.com",
		FavouriteColor: "#000-",
	}

	if err := e.Validator.Validate(user); err != nil {
        fmt.Println("Error when validate - ", err)
        panic(err)
    }
}
```
