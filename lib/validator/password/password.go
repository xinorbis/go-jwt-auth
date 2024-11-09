package password

import (
	"github.com/go-playground/validator/v10"
)

func Validate(fl validator.FieldLevel) bool {
	password := fl.Parent().Field(0).Field(1).String() //TODO: rewrite for any case
	passwordAgain := fl.Field().String()

	return password == passwordAgain
}
