package register

import (
	userReg "auth_service/internal/dto/user-reg"
	"auth_service/internal/services/register"
	"auth_service/lib/api/request"
	resp "auth_service/lib/api/response"
	notifierPkg "auth_service/lib/err-notifier/notifier"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

type Response struct {
	resp.Response
	UserID uint64 `json:"user-id,omitempty"`
}

func New(log *slog.Logger, registerService register.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.user.register.New"

		log = log.With(slog.String("operation", op))
		notifier := notifierPkg.New(w, r, log)
		registerService.SetNotifier(notifier)

		var user userReg.UserReg
		if err := request.Check(notifier, &user); err != nil {
			return
		}

		id, err := registerService.SaveUser(user.Email, user.Password)
		if err != nil {
			return
		}

		responseOK(w, r, id)
	}
}

func responseOK(w http.ResponseWriter, r *http.Request, userID uint64) {
	render.JSON(w, r, Response{
		Response: resp.OK(),
		UserID:   userID,
	})
}
