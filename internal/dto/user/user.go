package user

type User struct {
	ID       uint64
	Email    string
	Password string
	IsActive bool
}

type FullDataUser struct {
	ID       uint64
	Email    string
	Password string
	IsActive bool
	APIKey   string
}

type UserActive bool

const (
	IsActive    UserActive = true
	IsNotActive UserActive = false
)

func New(email, password string, isActive UserActive) *User {
	return &User{
		Email:    email,
		Password: password,
		IsActive: bool(isActive),
	}
}
