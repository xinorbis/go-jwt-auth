package revoke_api_key

type APIKey struct {
	APIKey string `json:"api-key" validate:"required"`
}

func (ak APIKey) GetLogErrorText() string {
	return "failed to decode request api-key body"
}

func (ak APIKey) GetResponseErrorText() string {
	return "failed to decode request api-key body"
}
