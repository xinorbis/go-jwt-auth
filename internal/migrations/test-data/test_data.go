package test_data

import (
	appDTO "auth_service/internal/dto/app"
	"gorm.io/gorm"
	"log"
)

var testApps = []*appDTO.App{
	{0, "Test", "test", "test"},
	{0, "Test2", "test2", "test2"},
}

type Migration struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *Migration {
	return &Migration{db}
}

func (m *Migration) MigrateUp() error {
	if !m.checkIfExists() {
		return m.DB.Create(testApps).Error
	}

	return nil
}

func (m *Migration) checkIfExists() bool {
	isExists := false

	var apps []appDTO.App
	var titles []string
	for _, app := range testApps {
		titles = append(titles, app.Title)
	}

	result := m.DB.Table("apps").Where("title IN ?", titles).Find(&apps)

	if result.Error != nil {
		panic(result.Error)
	}

	if len(apps) == len(testApps) {
		log.Println("migration already applied")
		isExists = true
	}

	return isExists
}
