package contorllers

import (
	"fmt"
	"net/http"
	"strconv"

	"profileyou/api/usecase"
	"profileyou/api/utils/errors"

	"github.com/gin-gonic/gin"
	// "gorm.io/driver/sqlite"
)

type keywordController struct {
	// keywordRepository repository.KeywordRepository
	keywordUseCase usecase.KeywordUseCase
}

// likes to Usecase by "ku"
func NewKeywordController(ku usecase.KeywordUseCase) keywordController {
	return keywordController{
		keywordUseCase: ku,
	}

}

func (ku *keywordController) GetAllKeywordsGin(c *gin.Context) {
	keywords, err := ku.keywordUseCase.GetKeywords()
	if err != nil {
		fmt.Println(err)
		apiErr := errors.NewBadRequestError("Bad Request")
		c.IndentedJSON(apiErr.Status, apiErr)
		return
	}
	c.IndentedJSON(http.StatusOK, keywords)
}

func (ku *keywordController) GetKeyword(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		apiErr := errors.NewBadRequestError("Bad Request")
		c.IndentedJSON(apiErr.Status, apiErr)
		return
	}

	keyword, err := ku.keywordUseCase.GetKeyword(id)
	if err != nil {
		fmt.Printf("Error %v", err)
		// c.JSON(http.StatusNotFound, errorResponse(err))
		apiErr := errors.NotFoundError("Not found")
		c.IndentedJSON(apiErr.Status, apiErr)
		return
	}

	c.IndentedJSON(http.StatusOK, keyword)

}

func (ku *keywordController) Index(c *gin.Context) {
	keywords, err := ku.keywordUseCase.GetKeywords()
	if err != nil {
		fmt.Println(err)
		apiErr := errors.NewBadRequestError("Bad Request")
		c.IndentedJSON(apiErr.Status, apiErr)
		return
	}
	c.IndentedJSON(http.StatusOK, keywords)
}

func (ku *keywordController) DetailKeyword(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		return
	}

	keyword, err := ku.keywordUseCase.GetKeyword(id)
	if err != nil {
		fmt.Println(err)
		apiErr := errors.NotFoundError("Not found")
		c.IndentedJSON(apiErr.Status, apiErr)
		return
	}
	c.IndentedJSON(http.StatusOK, keyword)
}

func (ku *keywordController) CreateKeyword(c *gin.Context) {
	word := c.Param("word")
	fmt.Printf("Receive a post: %s", word)

	// keyword := model.Keyword{Word: word}
	err := ku.keywordUseCase.CreateKeyword(word)
	if err != nil {
		fmt.Println(err)
		apiErr := errors.InternalSeverError("Server Error")
		c.IndentedJSON(apiErr.Status, apiErr)
		return
	}

	c.Redirect(301, "/")
}

func (ku *keywordController) UpdateKeyword(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Printf("Updating a keyword id: %d", id)
	keyword, err := ku.keywordUseCase.GetKeyword(id)
	if err != nil {
		fmt.Println(err)
		apiErr := errors.NotFoundError("Not found")
		c.IndentedJSON(apiErr.Status, apiErr)
		return
	}

	word := c.Param("word")
	description := c.Param("description")

	keyword.Word = word
	keyword.Description = description
	err = ku.keywordUseCase.UpdateKeyword(id, word, description)
	if err != nil {
		fmt.Println(err)
		apiErr := errors.InternalSeverError("Server Error")
		c.IndentedJSON(apiErr.Status, apiErr)
		return
	}

	c.IndentedJSON(http.StatusOK, keyword)
}

func (ku *keywordController) DeleteKeyword(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		apiErr := errors.NewBadRequestError("Bad request")
		c.IndentedJSON(apiErr.Status, apiErr)
	}

	keyword, err := ku.keywordUseCase.GetKeyword(id)
	if err != nil {
		fmt.Println(err)
		apiErr := errors.NotFoundError("Not found")
		c.IndentedJSON(apiErr.Status, apiErr)
	}

	err = ku.keywordUseCase.DeleteKeyword(id)
	if err != nil {
		fmt.Println(err)
		apiErr := errors.InternalSeverError("Server Error")
		c.IndentedJSON(apiErr.Status, apiErr)
		return
	}

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
