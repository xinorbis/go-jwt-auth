package db_err_notifier

import (
	"auth_service/internal/storage"
	errNotifier "auth_service/lib/err-notifier"
	"auth_service/lib/err-notifier/notifier"
)

type Notifier struct {
	errNotifier.ErrNotifier
}

func Notify(notifier notifier.Notifier, err error, logText, logErrorText, JSONText, JSONErrorText string) {
	if err.Error() == storage.DBErrNotFound {
		notifier.Notify(err, logText, logErrorText)
	} else {
		notifier.Notify(err, JSONText, JSONErrorText)
	}
}
