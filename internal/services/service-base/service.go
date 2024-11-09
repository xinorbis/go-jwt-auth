package service_base

import (
	"auth_service/lib/err-notifier/notifier"
)

type Service struct {
	notifier notifier.Notifier
}

func (s *Service) SetNotifier(notifier notifier.Notifier) {
	s.notifier = notifier
}

func (s *Service) GetNotifier() *notifier.Notifier {
	return &s.notifier
}
