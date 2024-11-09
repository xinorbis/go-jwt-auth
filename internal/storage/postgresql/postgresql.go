package postgresql

import (
	"auth_service/internal/config"
	testData "auth_service/internal/migrations/test-data"
	accessTokenModel "auth_service/internal/storage/models/access-token"
	apiKeyModelPkg "auth_service/internal/storage/models/api-key"
	appModelPkg "auth_service/internal/storage/models/app"
	refreshTokenModel "auth_service/internal/storage/models/refresh-token"
	userModelPkg "auth_service/internal/storage/models/user"
	accessTokenSchema "auth_service/internal/storage/schema/access-token"
	apiKeySchema "auth_service/internal/storage/schema/api-key"
	appSchema "auth_service/internal/storage/schema/app"
	refreshTokenSchema "auth_service/internal/storage/schema/refresh-token"
	userSchema "auth_service/internal/storage/schema/user"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	User         userModelPkg.User
	App          appModelPkg.App
	AccessToken  accessTokenModel.AccessToken
	RefreshToken refreshTokenModel.RefreshToken
	APIKey       apiKeyModelPkg.APIKey
}

func New(config config.DBConfig) (*Storage, error) {
	const op = "storage.postgresql.New"

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.Host, config.User, config.Password, config.DBName, config.Port, config.SSLMode, config.TimeZone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if config.UseORMMigrations {
		err = db.AutoMigrate(
			&userSchema.User{},
			&accessTokenSchema.AccessToken{},
			&refreshTokenSchema.RefreshToken{},
			&appSchema.App{},
			&apiKeySchema.APIKey{},
		)

		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
	}

	if config.MigrateTestData {
		migration := testData.New(db)
		err = migration.MigrateUp()
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
	}

	userModel := userModelPkg.New(db)
	appModel := appModelPkg.New(db)
	accessToken := accessTokenModel.New(db)
	refreshToken := refreshTokenModel.New(db)
	apiKey := apiKeyModelPkg.New(db)

	return &Storage{
		User:         userModel,
		App:          appModel,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		APIKey:       apiKey,
	}, nil
}
