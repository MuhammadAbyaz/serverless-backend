package interfaces

import "github.com/go-playground/validator/v10"

type DtoInterface interface {
	Validate(body any) error
}

type Validatable struct{}

func (v *Validatable) Validate(body any) error {
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		validationsErrr := err.(validator.ValidationErrors)
		return validationsErrr
	}
	return nil
}
