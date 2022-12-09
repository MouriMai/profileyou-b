package main

import (
	"net/http"
	"time"

	sqlite "profileyou/api/config/database"
	controllers "profileyou/api/controllers"
	"profileyou/api/infrastructure/persistance"
	"profileyou/internal/repository"
	"profileyou/internal/repository/dbrepo"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	// "gorm.io/driver/sqlite"
)

// const port = 8080

type application struct {
	DSN          string
	Domain       string
	DB           repository.DatabaseRepo
	JWTSecret    string
	JWTIssuer    string
	JWTAudience  string
	CookieDomain string
	APIKey       string
}

func main() {
	// set application config
	var app application

	// connect to the database
	db := sqlite.New()

	// Seed datas
	// db.Create(&models.Keyword{Word: "クリスマス", Description: "", ImageUrl: "test"})
	// db.Create(&models.Keyword{Word: "お正月", Description: "", ImageUrl: "test"})

	connect, err := db.DB()
	if err != nil {
		panic(err)
	}
	app.DB = &dbrepo.SQliteDBRepo{DB: connect}
	defer app.DB.Connection().Close()

	keywordRepository := persistance.NewKeywordPersistance(db)
	keywordController := controllers.NewKeywordController(keywordRepository)

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		// アクセスを許可したいアクセス元
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		// アクセスを許可したいHTTPメソッド(以下の例だとPUTやDELETEはアクセスできません)
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
			"DELETE",
		},
		// 許可したいHTTPリクエストヘッダ
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		// cookieなどの情報を必要とするかどうか
		AllowCredentials: true,
		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
	}))

	r.GET("/", controllers.GetAllKeywordsGin)
	// list all the keywords
	r.GET("/keywords", controllers.GetAllKeywordsGin)
	// list one keyword
	r.GET("/keywords/:id", controllers.GetKeyword)
	// create a new keyword
	r.POST("/keyword/create/:word", keywordController.CreateKeyword)
	r.POST("/keyword/update", keywordController.UpdateKeyword)
	r.DELETE("/keyword/delete/:id", keywordController.DeleteKeyword)
	r.GET("/message", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world",
		})
	})
	r.Run()

}
