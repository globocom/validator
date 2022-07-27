package validator

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestCPFValidation(t *testing.T) {
	type User struct {
		CPF string `validate:"cpf"`
	}

	userSuccess := &User{
		CPF: "273.487.260-97",
	}

	userError := &User{
		CPF: "273.487.260-11",
	}

	type fields struct {
		validator *validator.Validate
	}
	type args struct {
		testStruct interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "CPF validation success",
			fields:  fields{validator: New()},
			args:    args{userSuccess},
			wantErr: false,
		},
		{
			name:    "CPF validation error",
			fields:  fields{validator: New()},
			args:    args{userError},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.fields.validator.Struct(tt.args.testStruct); (err != nil) != tt.wantErr {
				t.Errorf("CPFValidation() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
