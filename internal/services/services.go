package services

import (
	APIKey "auth_service/internal/services/api-key"
	"auth_service/internal/services/app"
	"auth_service/internal/services/auth"
	"auth_service/internal/services/register"
	"auth_service/internal/services/token"
)

type Services struct {
	AuthService     auth.Service
	RegisterService register.Service
	APIKeyService   APIKey.Service
	AppService      app.Service
	AccessToken     token.Service
	RefreshToken    token.Service
}

func New(
	authService auth.Service,
	registerService register.Service,
	APIKeyService APIKey.Service,
	appService app.Service,
	accessToken token.Service,
	refreshToken token.Service,
) *Services {
	services := Services{
		AuthService:     authService,
		RegisterService: registerService,
		APIKeyService:   APIKeyService,
		AppService:      appService,
		AccessToken:     accessToken,
		RefreshToken:    refreshToken,
	}
	return &services
}
