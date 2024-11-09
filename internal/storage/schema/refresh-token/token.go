package refresh_token

import (
	"auth_service/internal/storage/schema/app"
	"auth_service/internal/storage/schema/user"
	"time"
)

type RefreshToken struct {
	ID        uint64    `gorm:"primaryKey;unique;autoIncrement:true"`
	Token     string    `gorm:"type:varchar(255);not null"`
	ExpiresAt time.Time `gorm:"not null"`
	UserID    uint64    `gorm:"column:user_id;not null;index:idx_unique_rt,unique"`
	User      user.User
	AppID     uint64 `gorm:"column:app_id;not null;index:idx_unique_rt,unique"`
	App       app.App
}
