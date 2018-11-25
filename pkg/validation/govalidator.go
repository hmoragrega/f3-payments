package validation

import (
	"github.com/asaskevich/govalidator"
)

// GoValidator uses thrid party govalidator to create a validator
type GoValidator struct{}

// Validate checks if the given parameter is valid
func (v *GoValidator) Validate(i interface{}) error {
	_, error := govalidator.ValidateStruct(i)

	return error
}
