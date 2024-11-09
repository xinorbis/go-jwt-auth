package api_key

import (
	"auth_service/internal/storage/schema/user"
)

type APIKey struct {
	ID     uint64 `gorm:"primaryKey;unique;autoIncrement:true"`
	Key    string `gorm:"type:varchar(255);unique;not null"`
	UserID uint64 `gorm:"column:user_id;unique;not null"`
	User   user.User
}
