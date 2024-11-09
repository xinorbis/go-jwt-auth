package refresh_access_token

type RespToken struct {
	RefreshToken string `json:"refresh-token" validate:"required"`
}

func (rt RespToken) GetLogErrorText() string {
	return "failed to decode token request body"
}

func (rt RespToken) GetResponseErrorText() string {
	return "failed to decode token request body"
}
