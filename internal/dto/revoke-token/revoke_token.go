package revoke_token

type Token struct {
	APIKey       string `json:"api-key" validate:"required"`
	RefreshToken string `json:"refresh-token" validate:"required"`
}

func (t Token) GetLogErrorText() string {
	return "failed to decode token request body"
}

func (t Token) GetResponseErrorText() string {
	return "failed to decode token request body"
}
