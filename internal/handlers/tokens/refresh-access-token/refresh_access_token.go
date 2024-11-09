package refresh_access_token

import (
	tokenDTO "auth_service/internal/dto/token"
	"auth_service/internal/services/app"
	"auth_service/internal/services/token"
	resp "auth_service/lib/api/response"
	notifierPkg "auth_service/lib/err-notifier/notifier"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

type Response struct {
	resp.Response
	AccessToken string `json:"access-token,omitempty"`
}

func New(log *slog.Logger, refreshTokenService token.Service, appService app.Service, accessTokenService token.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.tokens.refresh_access_token.New"

		log = log.With(slog.String("operation", op))
		notifier := notifierPkg.New(w, r, log)
		refreshTokenService.SetNotifier(notifier)

		var respTokenDTO tokenDTO.Token
		if err := refreshTokenService.CheckToken(r, &respTokenDTO); err != nil {
			return
		}

		appDTO, err := appService.Find(respTokenDTO.AppID)
		if err != nil {
			return
		}

		err = refreshTokenService.ValidateToken(respTokenDTO.Token, appDTO.Secret)
		if err != nil {
			return
		}

		newToken, err := accessTokenService.RegenAndSaveToken(respTokenDTO.UserID, appDTO)
		if err != nil {
			return
		}

		responseOK(w, r, newToken.Token)
	}
}

func responseOK(w http.ResponseWriter, r *http.Request, token string) {
	render.JSON(w, r, Response{
		Response:    resp.OK(),
		AccessToken: token,
	})
}
