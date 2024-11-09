package auth

import (
	APIKeyDTOPkg "auth_service/internal/dto/api-key"
	userDto "auth_service/internal/dto/user"
	userRequest "auth_service/internal/dto/user-request"
	serviceBase "auth_service/internal/services/service-base"
	"auth_service/internal/storage"
	APIKeyModel "auth_service/internal/storage/models/api-key"
	userModel "auth_service/internal/storage/models/user"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type Service struct {
	serviceBase.Service
	UserModel   userModel.User
	APIKeyModel APIKeyModel.APIKey
}

func New(userModel userModel.User, APIKey APIKeyModel.APIKey) Service {
	authService := Service{}
	authService.UserModel = userModel
	authService.APIKeyModel = APIKey

	return authService
}

func (s *Service) IsUserAuth(r *http.Request) (APIKeyDTOPkg.ApiKey, error) {
	APIKey := r.Header.Get("Bearer")
	notifier := s.GetNotifier()
	notifier.SetStatusUnauthorized()

	APIKeyDTO, err := s.APIKeyModel.GetKeyDataByKey(APIKey)
	if err != nil {
		notifier := s.GetNotifier()
		if err.Error() == storage.DBErrNotFound {
			notifier.Notify(err, "api key not found", storage.DBErrAPIKeyNotFound.Error())
		} else {
			notifier.Notify(err, "api key search error", storage.ErrInternalServerError.Error())
		}
	}

	return APIKeyDTO, err
}

func (s *Service) findUserByEmail(email string) (*userDto.FullDataUser, error) {
	foundUser, err := s.UserModel.FindByEmail(email)
	if err != nil {
		notifier := s.GetNotifier()
		if err.Error() == storage.DBErrNotFound {
			notifier.Notify(err, "user not found", storage.ErrUserNotFound.Error())
		} else {
			notifier.Notify(err, "user search error", storage.ErrInternalServerError.Error())
		}
	}

	return foundUser, err
}

func (s *Service) comparePasswords(sentPassword, userPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(sentPassword))
	if err != nil {
		notifier := s.GetNotifier()
		notifier.Notify(err, "invalid password", storage.ErrUserNotFound.Error())
	}

	return err
}

func (s *Service) AuthUser(user userRequest.User) (*userDto.FullDataUser, error) {
	foundUser, err := s.findUserByEmail(user.Email)
	if err != nil {
		return foundUser, err
	}

	err = s.comparePasswords(user.Password, foundUser.Password)

	return foundUser, err
}
