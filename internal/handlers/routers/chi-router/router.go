package chi_router

import (
	revokeApiKey "auth_service/internal/handlers/api-key/revoke-api-key"
	generateAccessToken "auth_service/internal/handlers/tokens/access-token/generate-access-token"
	refreshAccessToken "auth_service/internal/handlers/tokens/refresh-access-token"
	generateRefreshToken "auth_service/internal/handlers/tokens/refresh-token/generate-refresh-token"
	revokeRefreshToken "auth_service/internal/handlers/tokens/refresh-token/revoke-refresh-token"
	"auth_service/internal/handlers/user/auth"
	"auth_service/internal/handlers/user/register"
	"auth_service/internal/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
)

func New(log *slog.Logger, services *services.Services) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Post("/register", register.New(log, services.RegisterService))
	router.Post("/login", auth.New(log, services.AuthService, services.APIKeyService))
	router.Delete("/revoke-api-key", revokeApiKey.New(log, services.APIKeyService))
	router.Post("/generate-access-token", generateAccessToken.New(log, services.AuthService, services.AppService, services.AccessToken))
	router.Post("/generate-refresh-token", generateRefreshToken.New(log, services.AuthService, services.AppService, services.RefreshToken))
	router.Delete("/revoke-refresh-token", revokeRefreshToken.New(log, services.AuthService, services.RefreshToken))
	router.Get("/refresh-access-token", refreshAccessToken.New(log, services.RefreshToken, services.AppService, services.AccessToken))

	return router
}
