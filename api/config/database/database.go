package database

import (
	"github.com/gin-gonic/gin"
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

	// Seed datas
	// db.Create(&models.Keyword{Word: "クリスマス", Description: "", ImageUrl: "test"})
	// db.Create(&models.Keyword{Word: "お正月", Description: "", ImageUrl: "test"})

	db.AutoMigrate(&Keyword{})

	return db
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
