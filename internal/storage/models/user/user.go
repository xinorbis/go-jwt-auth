package user

import (
	userDto "auth_service/internal/dto/user"
	"gorm.io/gorm"
)

type User struct {
	db       *gorm.DB
	ID       uint64
	Email    string
	Password string
	IsActive bool
}

func New(db *gorm.DB) User {
	return User{
		db: db,
	}
}

func (u *User) Create(user *userDto.User) (uint64, error) {
	result := u.db.Create(&user)

	return user.ID, result.Error
}

func (u *User) FindByEmail(email string) (*userDto.FullDataUser, error) {
	var user userDto.FullDataUser
	fields := `
u.id id,
u.email email,
u.password password,
u.is_active is_active,
ak.key api_key
`
	result := u.db.Table("users u").
		Select(fields).
		Joins("LEFT JOIN api_keys ak ON u.id = ak.user_id").
		Where("email = ?", email).
		First(&user)

	return &user, result.Error
}
