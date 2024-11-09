package resp_token

type RespToken struct {
	AppCode string `json:"app-code" validate:"required"`
}

func (rt RespToken) GetLogErrorText() string {
	return "failed to decode token request body"
}

func (rt RespToken) GetResponseErrorText() string {
	return "failed to decode token request body"
}
