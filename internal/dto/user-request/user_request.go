package user_request

type User struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (u User) GetLogErrorText() string {
	return "failed to decode request body"
}

func (u User) GetResponseErrorText() string {
	return "failed to decode request body"
}
