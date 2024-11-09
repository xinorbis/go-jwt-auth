package refresh_token

import (
	appDTOPkg "auth_service/internal/dto/app"
	"auth_service/internal/dto/token"
	"gorm.io/gorm"
	"time"
)

type RefreshToken struct {
	db        *gorm.DB
	ID        uint64
	Token     string
	ExpiresAt time.Time
	AppID     uint64
	UserID    uint64
}

func New(db *gorm.DB) RefreshToken {
	return RefreshToken{
		db: db,
	}
}

func (rt RefreshToken) GetTokenByAppID(userId, appId uint64) (token.Token, error) {
	var tokenDTO token.Token

	result := rt.db.Model(rt).Where("user_id = ? AND app_id = ?", userId, appId).First(&tokenDTO)

	return tokenDTO, result.Error
}

func (rt RefreshToken) GetTokenDataByToken(tokenStr string) (token.Token, error) {
	var tokenDTO token.Token

	result := rt.db.Model(rt).Where("token = ?", tokenStr).First(&tokenDTO)

	return tokenDTO, result.Error
}

func (rt RefreshToken) Save(accessTokenDTO *token.Token) error {
	return rt.db.Table("refresh_tokens").Save(&accessTokenDTO).Error
}

func (rt RefreshToken) DeleteByToken(token string) (int64, error) {
	result := rt.db.Where("token = ?", token).Delete(&rt)

	return result.RowsAffected, result.Error
}

func (rt RefreshToken) GetAppByToken(token string) (appDTOPkg.App, error) {
	var appDTO appDTOPkg.App
	fields := `
app.id id,
app.code code,
app.secret secret,
app.title title
`
	result := rt.db.Table("refresh_tokens rt").
		Select(fields).
		Joins("LEFT JOIN apps app ON app.id = rt.app_id").
		Where("rt.token = ?", token).
		First(&appDTO)

	return appDTO, result.Error
}
