package conf

import "github.com/go-playground/validator"

var Validate *validator.Validate

func SetValidator() {
	Validate = validator.New()
}
