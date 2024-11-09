package err_notifier

import (
	"log/slog"
	"net/http"
)

type ErrNotifier struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	Logger         *slog.Logger
	statusCode     int
}

func (en *ErrNotifier) SetStatusOK() {
	en.statusCode = http.StatusOK
}

func (en *ErrNotifier) SetStatusUnauthorized() {
	en.statusCode = http.StatusUnauthorized
}

func (en *ErrNotifier) SetStatusInternalServerError() {
	en.statusCode = http.StatusInternalServerError
}

func (en *ErrNotifier) GetStatusCode() int {
	return en.statusCode
}
