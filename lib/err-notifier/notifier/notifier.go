package notifier

import (
	resp "auth_service/lib/api/response"
	errNotifier "auth_service/lib/err-notifier"
	"auth_service/lib/logger/sl"
	"log/slog"
	"net/http"
)

type Notifier struct {
	errNotifier.ErrNotifier
}

func New(w http.ResponseWriter, r *http.Request, log *slog.Logger) Notifier {
	notifier := Notifier{}
	notifier.ResponseWriter = w
	notifier.Request = r
	notifier.Logger = log
	notifier.SetStatusOK()

	return notifier
}

func (n *Notifier) Notify(err error, logTest, JSONText string) {
	n.Logger.Error(logTest, sl.Err(err))
	resp.JSONError(n.ResponseWriter, n.Request, JSONText, n.GetStatusCode())
}
