package api_key

import (
	APIKeyDTO "auth_service/internal/dto/api-key"
	serviceBase "auth_service/internal/services/service-base"
	"auth_service/internal/storage"
	APIKeyModel "auth_service/internal/storage/models/api-key"
	"auth_service/lib/generator"
	"errors"
)

type Service struct {
	serviceBase.Service
	APIKeyModel APIKeyModel.APIKey
}

func New(APIKeyModel APIKeyModel.APIKey) Service {
	APIKeyService := Service{}
	APIKeyService.APIKeyModel = APIKeyModel

	return APIKeyService
}

func (s *Service) MakeAndSaveAPIKey(userId uint64) (string, error) {
	newApiKey := generator.GenerateUniqueHash()
	apiKeyDto := APIKeyDTO.New(userId, newApiKey)

	err := s.APIKeyModel.Create(&apiKeyDto)
	if err != nil {
		notifier := s.GetNotifier()
		notifier.SetStatusInternalServerError()
		notifier.Notify(err, "failed to create api key", storage.ErrInternalServerError.Error())
	}

	return newApiKey, err
}

func (s *Service) DeleteByKey(key string) error {
	count, err := s.APIKeyModel.DeleteByKey(key)

	if err != nil {
		notifier := s.GetNotifier()
		notifier.SetStatusInternalServerError()
		notifier.Notify(err, "failed to delete API key", storage.ErrInternalServerError.Error())
		return err
	}

	if count == 0 {
		err = errors.New("API key not found")
		notifier := s.GetNotifier()
		notifier.Notify(err, "failed to delete API key", "failed to delete API key")
	}

	return err
}
