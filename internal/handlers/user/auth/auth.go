package auth

import (
	userRequest "auth_service/internal/dto/user-request"
	APIKey "auth_service/internal/services/api-key"
	"auth_service/internal/services/auth"
	"auth_service/lib/api/request"
	resp "auth_service/lib/api/response"
	notifierPkg "auth_service/lib/err-notifier/notifier"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

type Response struct {
	resp.Response
	APIKey string `json:"api-key,omitempty"`
}

func New(log *slog.Logger, authService auth.Service, APIKeyService APIKey.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.user.auth.New"

		log = log.With(slog.String("operation", op))
		notifier := notifierPkg.New(w, r, log)
		authService.SetNotifier(notifier)

		var user userRequest.User
		if err := request.Check(notifier, &user); err != nil {
			return
		}

		foundUser, err := authService.AuthUser(user)
		if err != nil {
			return
		}

		if len(foundUser.APIKey) == 0 {
			newApiKey, err := APIKeyService.MakeAndSaveAPIKey(foundUser.ID)
			if err != nil {
				return
			}

			foundUser.APIKey = newApiKey
		}

		responseOK(w, r, foundUser.APIKey)
	}
}

func responseOK(w http.ResponseWriter, r *http.Request, APIKey string) {
	render.JSON(w, r, Response{
		Response: resp.OK(),
		APIKey:   APIKey,
	})
}
