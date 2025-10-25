package validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validate *validator.Validate
}

func NewValidator() *Validator {
	v := &Validator{
		validate: validator.New(),
	}

	v.registerCustomerValidation()

	return v
}

func (v *Validator) Validate(i interface{}) (err error) {
	if err = v.validate.Struct(i); err != nil {
		return
	}

	return
}

func (v *Validator) registerCustomerValidation() {
	_ = v.validate.RegisterValidation("platenumber", validateIndonesianPlateNumber)
}

func validateIndonesianPlateNumber(fl validator.FieldLevel) bool {
	// Indonesian plate format: 1-2 letters + 1-4 digits + 1-3 letters
	// Examples: B1234VV, BA5678XXX, D9ABC
	pattern := `^[A-Z]{1,2}[0-9]{1,4}[A-Z]{1,3}$`
	matched, _ := regexp.MatchString(pattern, fl.Field().String())
	return matched
}
