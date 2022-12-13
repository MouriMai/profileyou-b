package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Keyword struct {
	gorm.Model
	KeywordId   string
	Word        string
	Description string
	ImageUrl    string
}

func New() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("mvc.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Seed datas
	// db.Create(Keyword{Word: "クリスマス", Description: "", ImageUrl: "test", KeywordId: "1"})
	// db.Create(Keyword{Word: "お正月", Description: "", ImageUrl: "test", KeywordId: "2"})

	db.AutoMigrate(&Keyword{})

	return db
}
