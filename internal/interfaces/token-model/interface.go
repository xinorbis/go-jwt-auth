package token_model

import (
	appDTOPkg "auth_service/internal/dto/app"
	"auth_service/internal/dto/token"
)

type Interface interface {
	GetTokenByAppID(userId, appId uint64) (token.Token, error)
	GetTokenDataByToken(tokenStr string) (token.Token, error)
	DeleteByToken(token string) (int64, error)
	GetAppByToken(token string) (appDTOPkg.App, error)
	Save(accessTokenDTO *token.Token) error
}
