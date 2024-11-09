package storage

import "errors"

var (
	ErrUserNotFound        = errors.New("user not found")
	ErrEmailExists         = errors.New("email already exists")
	ErrAppNotFound         = errors.New("app not found")
	ErrAPIKeyNotFound      = errors.New("API key not found")
	ErrAPIKeyGenerate      = errors.New("API key generation error")
	ErrInternalServerError = errors.New("internal server error") // что бы не показывать конкретики пользователю

	DBErrDuplicateEmail        = "ERROR: duplicate key value violates unique constraint \"uni_users_email\" (SQLSTATE 23505)"
	DBErrNotFound              = "record not found"
	DBErrAPIKeyNotFound        = errors.New("API key not found")
	DBErrAppNotFound           = errors.New("app not found")
	DBErrAccessTokenSearch     = errors.New("access token not found")
	DBErrRefreshTokenSearch    = errors.New("refresh token not found")
	DBErrAccessTokenGeneration = errors.New("access token generation error")
	DBErrAccessTokenSave       = errors.New("access token save error")
	DBErrRefreshTokenSave      = errors.New("refresh token save error")
	DBErrRefreshTokenDelete    = errors.New("refresh token delete error")
	DBErrUserSearch            = errors.New("user search error")
)
