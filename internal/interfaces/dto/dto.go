package dto

type DTO interface {
	GetLogErrorText() string
	GetResponseErrorText() string
}
