package token

import (
	"time"
)

type Token struct {
	ID        uint64
	Token     string
	ExpiresAt time.Time
	AppID     uint64
	UserID    uint64
}

func New(userID uint64, appID uint64, tokenDuration time.Duration, JWTToken string) Token {
	tokenDTO := Token{}
	tokenDTO.Token = JWTToken
	tokenDTO.UserID = userID
	tokenDTO.AppID = appID
	tokenDTO.ExpiresAt = time.Now().Add(tokenDuration)

	return tokenDTO
}

func (t *Token) Load(token Token) {
	t.Token = token.Token
	t.UserID = token.UserID
	t.AppID = token.AppID
	t.ExpiresAt = token.ExpiresAt
}
