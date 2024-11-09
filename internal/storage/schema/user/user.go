package user

import (
	userDto "auth_service/internal/dto/user"
	"gorm.io/gorm"
)

type User struct {
	ID       uint64 `gorm:"primaryKey;unique;autoIncrement:true"`
	Email    string `gorm:"type:varchar(64);unique;not null"`
	Password string `gorm:"type:varchar(128);not null"`
	IsActive bool   `gorm:"default:true;not null"`
}

func (u *User) Create(user userDto.User, db *gorm.DB) (uint64, error) {
	result := db.Create(&user)

	return user.ID, result.Error
}
