package generate_access_token

import (
	respToken "auth_service/internal/dto/resp-token"
	"auth_service/internal/services/app"
	"auth_service/internal/services/auth"
	"auth_service/internal/services/token"
	"auth_service/lib/api/request"
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

func New(log *slog.Logger, authService auth.Service, appService app.Service, accessTokenService token.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.tokens.generate_access_token.New"

		log = log.With(slog.String("operation", op))
		notifier := notifierPkg.New(w, r, log)
		authService.SetNotifier(notifier)

		APIKeyDTO, err := authService.IsUserAuth(r)
		if err != nil {
			return
		}

		var respTokenDTO respToken.RespToken
		if err := request.Check(notifier, &respTokenDTO); err != nil {
			return
		}

		appService.SetNotifier(notifier)
		appDTO, err := appService.FindAppByCode(respTokenDTO.AppCode)
		if err != nil {
			return
		}

		newToken, err := accessTokenService.RegenAndSaveToken(APIKeyDTO.UserID, appDTO)
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
