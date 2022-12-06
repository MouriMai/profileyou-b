package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Keyword struct {
	gorm.Model
	Word        string
	Description string
	ImageUrl    string
}

func New() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("mvc.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Keyword{})

	return db
}
