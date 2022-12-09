package contorllers

import (
	"fmt"
	"net/http"
	"strconv"

	sqlite "profileyou/api/config/database"
	"profileyou/api/domain/model"
	"profileyou/api/domain/repository"
	models "profileyou/api/models"

	"github.com/gin-gonic/gin"
	// "gorm.io/driver/sqlite"
)

type keywordController struct {
	keywordRepository repository.KeywordRepository
}

func GetAllKeywordsGin(c *gin.Context) {
	db := sqlite.New()

	connect, err := db.DB()
	if err != nil {
		panic(err)
	}

	var keywords []models.Keyword
	db.Find(&keywords)
	// fmt.Println(&keywords)
	connect.Close()

	if err != nil {
		fmt.Println(err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, keywords)
}

func GetKeyword(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	db := sqlite.New()

	connect, err := db.DB()
	if err != nil {
		panic(err)
	}

	var keyword models.Keyword
	db.First(&keyword, id)

	connect.Close()

	c.IndentedJSON(http.StatusOK, keyword)

}

func NewKeywordController(kr repository.KeywordRepository) keywordController {
	return keywordController{
		keywordRepository: kr,
	}

}

func (ku *keywordController) Index(c *gin.Context) {
	keywords := ku.keywordRepository.GetKeywords()
	c.HTML(200, "index.html", gin.H{"keywords": keywords})
}

func (ku *keywordController) DetailKeyword(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	keyword := ku.keywordRepository.GetKeyword(id)
	c.HTML(200, "detail.html", gin.H{"keyword": keyword})
}

func (ku *keywordController) CreateKeyword(c *gin.Context) {
	word := c.Param("word")
	fmt.Printf("Receive a post: %s", word)
	// age, _ := strconv.Atoi(c.PostForm("age"))

	keyword := model.Keyword{Word: word}
	ku.keywordRepository.Create(keyword)

	c.Redirect(301, "/")
}

func (ku *keywordController) UpdateKeyword(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Printf("Updating a keyword id: %d", id)
	keyword := ku.keywordRepository.GetKeyword(id)

	word := c.Param("word")
	description := c.Param("description")
	// age, _ := strconv.Atoi(c.PostForm("age"))

	keyword.Word = word
	keyword.Description = description
	ku.keywordRepository.Update(keyword)

	c.IndentedJSON(http.StatusOK, keyword)
	// c.Redirect(301, "/")
}

func (ku *keywordController) DeleteKeyword(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Printf("Deleting a keyword id: %d", id)
	keyword := ku.keywordRepository.GetKeyword(id)

	ku.keywordRepository.Delete(keyword)

	c.IndentedJSON(http.StatusOK, keyword)

	// c.Redirect(302, "/")
}

// func (app *application) GetAllKeywords(w http.ResponseWriter, r *http.Request) {
// 	keywords, err := app.DB.AllKeywords()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	out, err := json.Marshal(keywords)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(out)
// var keywords []models.Keyword

// Temtative list
// one := models.Keyword{
// 	ID:          1,
// 	Word:        "クリスマス",
// 	ImageUrl:    "",
// 	Description: "12月にプレゼント渡す",
// 	CreatedAt:   time.Now(),
// }
// keywords = append(keywords, one)

// two := models.Keyword{
// 	ID:          2,
// 	Word:        "テスラ",
// 	ImageUrl:    "",
// 	Description: "イーロンマスクがTwitter買収",
// 	CreatedAt:   time.Now(),
// }
// keywords = append(keywords, two)

// out, err := json.Marshal(keywords)
// if err != nil {
// 	fmt.Println(err)
// }

// c.JSONP(http.StatusOK, gin.H{
// 	"message": "ok",
// 	"data":    out,
// 	"lists":   keywords,
// })

// }

// func (app *application) GetKeyword(w http.ResponseWriter, r *http.Request) {
// 	// db := sqlite.New()

// 	// id := c.Param("id")
// 	id := chi.URLParam(r, "id")
// 	keywordID, err := strconv.Atoi(id)
// 	if err != nil {
// 		app.errorJSON(w, err)
// 		return
// 	}

// 	keyword, err := app.DB.OneMovie(keywordID)
// 	if err != nil {
// 		app.errorJSON(w, err)
// 		return
// 	}

// 	_ = app.writeJSON(w, http.StatusOK, keyword)

// }

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
