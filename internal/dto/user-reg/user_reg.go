package user_reg

import userRequest "auth_service/internal/dto/user-request"

type UserReg struct {
	userRequest.User
	PasswordAgain string `json:"password-again" validate:"required,password"`
}
