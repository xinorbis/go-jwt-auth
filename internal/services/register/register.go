package register

import (
	userDTOPkg "auth_service/internal/dto/user"
	serviceBase "auth_service/internal/services/service-base"
	"auth_service/internal/storage"
	userModel "auth_service/internal/storage/models/user"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	serviceBase.Service
	UserModel userModel.User
}

func New(userModel userModel.User) Service {
	registerService := Service{}
	registerService.UserModel = userModel

	return registerService
}

func (s *Service) encryptPassword(password string) (string, error) {
	passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		notifier := s.GetNotifier()
		notifier.SetStatusInternalServerError()
		notifier.Notify(err, "failed to generate password hash", storage.ErrInternalServerError.Error())
	}

	return string(passHash), err
}

func (s *Service) SaveUser(email, pass string) (uint64, error) {
	passHash, err := s.encryptPassword(pass)
	if err != nil {
		return 0, err
	}

	userDTO := userDTOPkg.New(email, passHash, userDTOPkg.IsActive)
	id, err := s.UserModel.Create(userDTO)

	if err != nil {
		notifier := s.GetNotifier()
		if err.Error() == storage.DBErrDuplicateEmail {
			notifier.Notify(err, "failed to create user", storage.ErrEmailExists.Error())
		} else {
			notifier.SetStatusInternalServerError()
			notifier.Notify(err, "failed to create user", storage.ErrInternalServerError.Error())
		}
	}

	return id, err
}
