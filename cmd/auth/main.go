package main

import (
	"auth_service/internal/config"
	chiRouter "auth_service/internal/handlers/routers/chi-router"
	servicesPkg "auth_service/internal/services"
	APIKey "auth_service/internal/services/api-key"
	"auth_service/internal/services/app"
	authService "auth_service/internal/services/auth"
	registerService "auth_service/internal/services/register"
	tokenService "auth_service/internal/services/token"
	"auth_service/internal/storage/postgresql"
	"auth_service/lib/logger/sl"
	"log/slog"
	"net/http"
	"os"
)

const (
	envLocal = "local"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()
	log := setupLogger(cfg.Env)

	storage, err := postgresql.New(cfg.DB)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	services := servicesPkg.New(
		authService.New(storage.User, storage.APIKey),
		registerService.New(storage.User),
		APIKey.New(storage.APIKey),
		app.New(storage.App),
		tokenService.New(storage.AccessToken, cfg.AccessTokenTTL),
		tokenService.New(storage.RefreshToken, cfg.RefreshTokenTTL),
	)

	router := chiRouter.New(log, services)

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start server: ", err)
	}

	log.Error("failed to start server")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
