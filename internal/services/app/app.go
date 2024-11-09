package app

import (
	appDTOPkg "auth_service/internal/dto/app"
	serviceBase "auth_service/internal/services/service-base"
	"auth_service/internal/storage"
	"auth_service/internal/storage/models/app"
)

type Service struct {
	serviceBase.Service
	AppModel app.App
}

func New(AppModel app.App) Service {
	appService := Service{}
	appService.AppModel = AppModel

	return appService
}

func (s Service) FindAppByCode(code string) (appDTOPkg.App, error) {
	appDTO, err := s.AppModel.FindByCode(code)
	if err != nil {
		notifier := s.GetNotifier()
		if err.Error() == storage.DBErrNotFound {
			notifier.Notify(err, "failed to find app code", storage.DBErrAppNotFound.Error())
		} else {
			notifier.SetStatusInternalServerError()
			notifier.Notify(err, "failed to find app code", storage.ErrInternalServerError.Error())
		}

		return appDTO, err
	}

	return appDTO, nil
}

func (s Service) Find(appID uint64) (appDTOPkg.App, error) {
	appDTO, err := s.AppModel.Find(appID)
	if err != nil {
		notifier := s.GetNotifier()
		if err.Error() == storage.DBErrNotFound {
			notifier.Notify(err, "failed to find app", storage.DBErrAppNotFound.Error())
		} else {
			notifier.SetStatusInternalServerError()
			notifier.Notify(err, "app search error", storage.ErrInternalServerError.Error())
		}

		return appDTOPkg.App{}, err
	}

	return appDTO, nil
}
