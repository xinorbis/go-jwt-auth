package request

import (
	"auth_service/internal/interfaces/dto"
	resp "auth_service/lib/api/response"
	"auth_service/lib/err-notifier/notifier"
	"auth_service/lib/validator/password"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

func DecodeJSON(notifier notifier.Notifier, dto dto.DTO) error {
	err := render.DecodeJSON(notifier.Request.Body, &dto)

	if err != nil {
		notifier.Notify(err, dto.GetLogErrorText(), dto.GetResponseErrorText())
	}

	return err
}

func Validate(notifier notifier.Notifier, dto dto.DTO) error {
	validate := validator.New()

	err := validate.RegisterValidation("password", password.Validate)
	if err != nil {
		validateErr := err.(validator.ValidationErrors)
		notifier.Notify(err, "password validation error", resp.ValidationError(validateErr).Error)

		return err
	}

	err = validate.Struct(dto)

	if err != nil {
		validateErr := err.(validator.ValidationErrors)
		notifier.Notify(err, "invalid request", resp.ValidationError(validateErr).Error)
	}

	return err
}

func Check(notifier notifier.Notifier, dto dto.DTO) error {
	err := DecodeJSON(notifier, dto)
	if err != nil {
		return err
	}

	err = Validate(notifier, dto)

	return err
}
