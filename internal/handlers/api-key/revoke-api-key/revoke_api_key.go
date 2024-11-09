package revoke_api_key

import (
	revokeAPIKey "auth_service/internal/dto/revoke-api-key"
	APIKey "auth_service/internal/services/api-key"
	"auth_service/lib/api/request"
	resp "auth_service/lib/api/response"
	notifierPkg "auth_service/lib/err-notifier/notifier"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

func New(log *slog.Logger, APIKeyService APIKey.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.revoke.revoke_api_key.New"

		log = log.With(slog.String("operation", op))
		notifier := notifierPkg.New(w, r, log)

		var sentAPIKey revokeAPIKey.APIKey
		if err := request.Check(notifier, &sentAPIKey); err != nil {
			return
		}

		APIKeyService.SetNotifier(notifier)
		if err := APIKeyService.DeleteByKey(sentAPIKey.APIKey); err != nil {
			return
		}

		responseOK(w, r)
	}
}

func responseOK(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, resp.OK())
}
