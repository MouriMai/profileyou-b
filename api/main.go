package main

import (
	"net/http"
	"time"

	controller "profileyou/api/controllers"
	"profileyou/internal/repository"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const port = 8080

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

func newWord(c *gin.Context) {

}

func main() {
	// set application config
	// var app application

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

	r.GET("/message", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world",
		})
	})

	// list all the images
	r.GET("/keywords", controller.GetAllList)
	r.GET("/keywords/:id", controller.GetKeyword)
	// r.GET("/keywords", app.AllKeywords)
	r.POST("/new", newWord)
	r.Run()
	// fmt.Println("Hello world.")

}
