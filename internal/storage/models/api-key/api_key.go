package api_key

import (
	apiKey "auth_service/internal/dto/api-key"
	"gorm.io/gorm"
)

type APIKey struct {
	db     *gorm.DB
	ID     uint64
	Key    string
	UserID uint64
}

func New(db *gorm.DB) APIKey {
	return APIKey{
		db: db,
	}
}

func (ak *APIKey) Create(APIKey *apiKey.ApiKey) error {
	return ak.db.Create(&APIKey).Error
}

func (ak *APIKey) DeleteByKey(key string) (int64, error) {
	result := ak.db.Where("key = ?", key).Delete(&ak)

	return result.RowsAffected, result.Error
}

func (ak *APIKey) GetKeyDataByKey(key string) (apiKey.ApiKey, error) {
	var apiKeyDto apiKey.ApiKey
	result := ak.db.Model(ak).Where("key = ?", key).First(&apiKeyDto)

	return apiKeyDto, result.Error
}
