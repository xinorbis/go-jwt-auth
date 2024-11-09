package app

type App struct {
	ID     uint64 `gorm:"primary_key;unique;auto_increment:true"`
	Title  string `gorm:"type:varchar(64);unique;not null"`
	Code   string `gorm:"type:varchar(32);unique;not null"`
	Secret string `gorm:"type:varchar(64);unique;not null"`
}
