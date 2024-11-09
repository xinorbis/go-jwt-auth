package revoke_refresh_token

import (
	revokeToken "auth_service/internal/dto/revoke-token"
	"auth_service/internal/services/auth"
	"auth_service/internal/services/token"
	"auth_service/lib/api/request"
	resp "auth_service/lib/api/response"
	notifierPkg "auth_service/lib/err-notifier/notifier"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

func New(log *slog.Logger, authService auth.Service, refreshTokenService token.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.revoke.revoke_refresh_token.New"

		log = log.With(slog.String("operation", op))
		notifier := notifierPkg.New(w, r, log)
		authService.SetNotifier(notifier)

		_, err := authService.IsUserAuth(r)
		if err != nil {
			return
		}

		var revokeTokenDTO revokeToken.Token
		if err := request.Check(notifier, &revokeTokenDTO); err != nil {
			return
		}

		refreshTokenService.SetNotifier(notifier)
		if err := refreshTokenService.DeleteToken(revokeTokenDTO.RefreshToken); err != nil {
			return
		}

		responseOK(w, r)
	}
}

func responseOK(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, resp.OK())
}
