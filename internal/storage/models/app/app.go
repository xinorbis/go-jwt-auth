package app

import (
	appDTO "auth_service/internal/dto/app"
	"gorm.io/gorm"
)

type App struct {
	db    *gorm.DB
	ID    uint64
	Title string
	Code  string
}

func New(db *gorm.DB) App {
	return App{
		db: db,
	}
}

func (a *App) FindByCode(appTitle string) (appDTO.App, error) {
	var app appDTO.App
	result := a.db.Where("code = ?", appTitle).First(&app)

	return app, result.Error
}

func (a *App) Find(appID uint64) (appDTO.App, error) {
	var app appDTO.App
	result := a.db.Model(a).First(&app, appID)

	return app, result.Error
}
