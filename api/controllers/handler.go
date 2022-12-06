package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"profileyou/api/models"
	"strconv"
	"time"

	sqlite "profileyou/api/config/database"

	"github.com/gin-gonic/gin"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world")
}

// // AllMovies returns a slice of all movies as JSON.
// func (app *application) AllKeywords(c *gin.Context) {
// 	keywords, err := app.DB.AllMovies()
// 	if err != nil {
// 		app.errorJSON(c, err)
// 		return
// 	}

// 	_ = app.writeJSON(c, http.StatusOK, keywords)
// }

func GetAllList(c *gin.Context) {
	var lists []models.Keyword

	// Temtative list
	one := models.Keyword{
		ID:          1,
		Word:        "クリスマス",
		ImageUrl:    "",
		Description: "12月にプレゼント渡す",
		CreatedAt:   time.Now(),
	}
	lists = append(lists, one)

	two := models.Keyword{
		ID:          2,
		Word:        "テスラ",
		ImageUrl:    "",
		Description: "イーロンマスクがTwitter買収",
		CreatedAt:   time.Now(),
	}
	lists = append(lists, two)

	out, err := json.Marshal(lists)
	if err != nil {
		fmt.Println(err)
	}

	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
		"data":    out,
		"lists":   lists,
	})

}

func GetKeyword(c *gin.Context) {
	db := sqlite.New()

	id := c.Param("id")
	keywordID, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}

	keyword := model.GetCustomer(keywordID)
	// c.HTML(200, "detail.html", gin.H{"keyword": keyword})

	return

}

// func GetKeywords() []Keyword {
// 	db := sqlite.New()

// 	connect, err := db.DB()
// 	if err != nil {
// 		panic(err)
// 	}

// 	var keywords []Keyword
// 	db.Find(&keywords)

// 	connect.Close()

// 	return keywords
// }

// func (c *Keyword) Create() {
// 	db := sqlite.New()

// 	connect, err := db.DB()
// 	if err != nil {
// 		panic(err)
// 	}

// 	db.Create(c)

// 	connect.Close()
// }

// func (c *Keyword) Update() {
// 	db := sqlite.New()

// 	connect, err := db.DB()
// 	if err != nil {
// 		panic(err)
// 	}

// 	db.Save(c)

// 	connect.Close()
// }

// func (c *Keyword) Delete() {
// 	db := sqlite.New()

// 	connect, err := db.DB()
// 	if err != nil {
// 		panic(err)
// 	}

// 	db.Delete(c)

// 	connect.Close()
// }
