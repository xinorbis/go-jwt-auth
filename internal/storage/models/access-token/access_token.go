package access_token

import (
	appDTOPkg "auth_service/internal/dto/app"
	"auth_service/internal/dto/token"
	"gorm.io/gorm"
	"time"
)

type AccessToken struct {
	db        *gorm.DB
	ID        uint64
	Token     string
	ExpiresAt time.Time
	AppID     uint64
	UserID    uint64
}

func New(db *gorm.DB) AccessToken {
	return AccessToken{
		db: db,
	}
}

func (at AccessToken) GetTokenByAppID(userId, appId uint64) (token.Token, error) {
	var tokenDTO token.Token
	result := at.db.Model(at).Where("user_id = ? AND app_id = ?", userId, appId).First(&tokenDTO)

	return tokenDTO, result.Error
}

func (at AccessToken) GetTokenDataByToken(tokenStr string) (token.Token, error) {
	var tokenDTO token.Token

	result := at.db.Model(at).Where("token = ?", tokenStr).First(&tokenDTO)

	return tokenDTO, result.Error
}

func (at AccessToken) Save(accessTokenDTO *token.Token) error {
	return at.db.Table("access_tokens").Save(&accessTokenDTO).Error
}

func (at AccessToken) DeleteByToken(token string) (int64, error) {
	result := at.db.Where("token = ?", token).Delete(&at)

	return result.RowsAffected, result.Error
}

func (at AccessToken) GetAppByToken(token string) (appDTOPkg.App, error) {
	var appDTO appDTOPkg.App
	fields := `
app.id id,
app.code code,
app.secret secret,
app.title title
`
	result := at.db.Table("access_tokens rt").
		Select(fields).
		Joins("LEFT JOIN apps app ON app.id = rt.app_id").
		Where("rt.token = ?", token).
		First(&appDTO)

	return appDTO, result.Error
}
