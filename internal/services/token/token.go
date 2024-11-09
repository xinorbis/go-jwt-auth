package token

import (
	appDTO "auth_service/internal/dto/app"
	"auth_service/internal/dto/token"
	tokenModel "auth_service/internal/interfaces/token-model"
	serviceBase "auth_service/internal/services/service-base"
	"auth_service/internal/storage"
	"auth_service/lib/jwt"
	"errors"
	"net/http"
	"time"
)

type Service struct {
	serviceBase.Service
	TokenModel    tokenModel.Interface
	tokenDuration time.Duration
}

func New(tokenModel tokenModel.Interface, tokenDuration time.Duration) Service {
	tokenService := Service{}
	tokenService.TokenModel = tokenModel
	tokenService.tokenDuration = tokenDuration

	return tokenService
}

func (s *Service) GetTokenByAppID(userID, appID uint64) (token.Token, error) {
	userAccessToken, err := s.TokenModel.GetTokenByAppID(userID, appID)

	if err != nil && err.Error() != storage.DBErrNotFound {
		notifier := s.GetNotifier()
		notifier.SetStatusInternalServerError()
		notifier.Notify(err, "failed to get access token", storage.ErrInternalServerError.Error())

		return token.Token{}, err
	}

	return userAccessToken, nil
}

func (s *Service) generateJWT(userID uint64, app appDTO.App) (token.Token, error) {
	tokenDTO := token.Token{}

	tokenExpDate := time.Now().Add(s.tokenDuration).Unix()
	tokenStr, err := jwt.NewToken(userID, app.Code, app.Secret, tokenExpDate)
	if err != nil {
		notifier := s.GetNotifier()
		notifier.Notify(err, "failed to generate access token", storage.DBErrAccessTokenGeneration.Error())

		return tokenDTO, err
	}

	tokenDTO = token.New(userID, app.ID, s.tokenDuration, tokenStr)

	return tokenDTO, err
}

func (s *Service) saveToken(tokenDTO, userAccessToken token.Token) (token.Token, error) {
	userAccessToken.Load(tokenDTO)

	err := s.TokenModel.Save(&userAccessToken)
	if err != nil {
		notifier := s.GetNotifier()
		notifier.SetStatusInternalServerError()
		notifier.Notify(err, "failed to save access token", storage.ErrInternalServerError.Error())
	}

	return userAccessToken, err
}

func (s *Service) RegenAndSaveToken(userID uint64, app appDTO.App) (token.Token, error) {
	foundTokenDTO, err := s.GetTokenByAppID(userID, app.ID)
	if err != nil {
		return token.Token{}, err
	}

	generatedTokenDTO, err := s.generateJWT(userID, app)
	if err != nil {
		return token.Token{}, err
	}

	newToken, err := s.saveToken(generatedTokenDTO, foundTokenDTO)
	if err != nil {
		return token.Token{}, err
	}

	return newToken, nil
}

func (s *Service) GetTokenDataByToken(token string) (token.Token, error) {
	userRefreshToken, err := s.TokenModel.GetTokenDataByToken(token)
	if err != nil {
		notifier := s.GetNotifier()
		if err.Error() == storage.DBErrNotFound {
			notifier.Notify(err, "failed to get refresh token", storage.DBErrRefreshTokenSearch.Error())
		} else {
			notifier.SetStatusInternalServerError()
			notifier.Notify(err, "refresh token search error", storage.ErrInternalServerError.Error())
		}
	}

	return userRefreshToken, err
}

func (s *Service) deleteByToken(token string) error {
	count, err := s.TokenModel.DeleteByToken(token)

	if err != nil {
		notifier := s.GetNotifier()
		notifier.SetStatusInternalServerError()
		notifier.Notify(err, "refresh token delete error", storage.ErrInternalServerError.Error())
		return err
	}

	if count == 0 {
		err = errors.New("refresh token not found")
		notifier := s.GetNotifier()
		notifier.Notify(err, "failed to delete refresh token", storage.DBErrRefreshTokenDelete.Error())
	}

	return err
}

func (s *Service) DeleteToken(token string) error {
	_, err := s.GetTokenDataByToken(token)
	if err != nil {
		return err
	}

	return s.deleteByToken(token)
}

func (s *Service) ValidateToken(token, secret string) error {
	err := jwt.ParseToken(token, secret)

	if err != nil {
		errText := "invalid refresh token"
		notifier := s.GetNotifier()
		notifier.Notify(err, "failed to validate refresh token", errText)

		return err
	}

	return nil
}

func (s *Service) CheckToken(r *http.Request, tokenDTO *token.Token) error {
	bearerToken := r.Header.Get("Bearer")

	userRefreshToken, err := s.GetTokenDataByToken(bearerToken)
	if err != nil {
		return err
	}

	tokenDTO.Load(userRefreshToken)

	return nil
}
