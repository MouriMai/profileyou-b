package contorllers

import (
	"fmt"
	"net/http"

	"profileyou/api/usecase"
	"profileyou/api/utils/errors"

	"github.com/gin-gonic/gin"
	// "gorm.io/driver/sqlite"
)

type keywordController struct {
	keywordUseCase usecase.KeywordUseCase
}

// likes to Usecase by "ku"
func NewKeywordController(ku usecase.KeywordUseCase) keywordController {
	return keywordController{
		keywordUseCase: ku,
	}

}

func (kc *keywordController) GetAllKeywordsGin(c *gin.Context) {
	// fmt.Println("GET ALL KEYWORDS")
	// keywords, err := ku.keywordUseCase.GetKeywords()
	// fmt.Printf("RETRIEVE KEYWORDS, %v", keywords)
	// if err != nil {
	// 	fmt.Println(err)
	// 	apiErr := errors.NewBadRequestError("Bad Request")
	// 	c.IndentedJSON(apiErr.Status, apiErr)
	// 	return
	// }
	// c.IndentedJSON(http.StatusOK, keywords)
	fmt.Println("GET ALL KEYWORDS")
	keywords, err := kc.keywordUseCase.GetKeywords()
	fmt.Printf("HANDLER: RETRIEVE KEYWORDS, %v\n", keywords)
	if err != nil {
		fmt.Println(err)
		apiErr := errors.NewBadRequestError("Bad Request")
		c.IndentedJSON(apiErr.Status, apiErr)
		return
	}
	// 20221213 - Deleted prior to implement ResultDataField struct
	// c.IndentedJSON(http.StatusOK, keywords)

	type ResultDataField struct {
		KeywordId   string
		Word        string
		Description string
		ImageUrl    string
	}
	var data []ResultDataField
	for _, keyword := range keywords {
		keywordId := string(keyword.GetKeywordId())
		word := string(keyword.GetWord())
		description := string(keyword.GetDescription())
		imageUrl := string(keyword.GetImageUrl())
		data = append(data, ResultDataField{KeywordId: keywordId, Word: word, Description: description, ImageUrl: imageUrl})
	}
	c.IndentedJSON(http.StatusOK, data)
}

func (ku *keywordController) GetKeyword(c *gin.Context) {
	// id, err := strconv.Atoi(c.Param("id"))
	// if err != nil {
	// 	apiErr := errors.NewBadRequestError("Bad Request")
	// 	c.IndentedJSON(apiErr.Status, apiErr)
	// 	return
	// }
	id := c.Param("id")
	keyword, err := ku.keywordUseCase.GetKeyword(id)

	if err != nil {
		fmt.Printf("Error %v", err)
		apiErr := errors.NotFoundError("Not found")
		c.IndentedJSON(apiErr.Status, apiErr)
		return
	}

	c.IndentedJSON(http.StatusOK, keyword)

}

func (ku *keywordController) Index(c *gin.Context) {
	fmt.Println("GET ALL KEYWORDS")
	keywords, err := ku.keywordUseCase.GetKeywords()
	fmt.Printf("RETRIEVE KEYWORDS, %v", keywords)
	if err != nil {
		fmt.Println(err)
		apiErr := errors.NewBadRequestError("Bad Request")
		c.IndentedJSON(apiErr.Status, apiErr)
		return
	}
	// 20221213 - Deleted prior to implement ResultDataField struct
	// c.IndentedJSON(http.StatusOK, keywords)

	type ResultDataField struct {
		KeywordId   string
		Word        string
		Description string
	}
	var data []ResultDataField
	for _, keyword := range keywords {
		keywordId := string(keyword.GetKeywordId())
		word := string(keyword.GetWord())
		description := string(keyword.GetDescription())
		data = append(data, ResultDataField{KeywordId: keywordId, Word: word, Description: description})
	}
	c.IndentedJSON(http.StatusOK, data)
}

func (ku *keywordController) DetailKeyword(c *gin.Context) {
	// id, err := strconv.Atoi(c.Param("id"))
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// 20221213 - changed to id type from int to string
	// Getkeyword method's params also changed to string
	id := c.Param("id")
	keyword, err := ku.keywordUseCase.GetKeyword(id)
	if err != nil {
		fmt.Println(err)
		apiErr := errors.NotFoundError("Not found")
		c.IndentedJSON(apiErr.Status, apiErr)
		return
	}
	// c.IndentedJSON(http.StatusOK, keyword)

	type ResultDataField struct {
		KeywordId   string
		Word        string
		Description string
	}

	data := ResultDataField{KeywordId: string(keyword.GetKeywordId()), Word: string(keyword.GetWord()), Description: string(keyword.GetDescription())}
	c.IndentedJSON(http.StatusOK, data)
}

func (ku *keywordController) CreateKeyword(c *gin.Context) {

	// 20221213 - Validation for create
	// type RequestDataField struct {
	// 	Word string `form:"word" binding:"required"`
	// 	Description string `form:"description"`
	// 	ImageUrl string `form:"image_url"`
	// }

	// var form RequestDataField
	// if err := c.ShouldBind(&form); err != nil {
	// 	fmt.Println(err)
	// 	apiErr := errors.NewBadRequestError("Bad request")
	// 	c.IndentedJSON(apiErr.Status, apiErr)
	// 	return
	// }

	// word := form.Word

	//

	word := c.Param("word")
	description := c.Param("description")
	imageUrl := c.Param("imageUrl")
	fmt.Printf("Receive a post: %s", word)

	// keyword := model.Keyword{Word: word}
	err := ku.keywordUseCase.CreateKeyword(word, description, imageUrl)
	if err != nil {
		fmt.Println(err)
		apiErr := errors.InternalSeverError("Server Error")
		c.IndentedJSON(apiErr.Status, apiErr)
		return
	}

	c.Redirect(301, "/")
}

func (ku *keywordController) UpdateKeyword(c *gin.Context) {
	type RequestDataField struct {
		ID          string `form:"id" binding:"required"`
		Word        string `form:"word" binding:"required"`
		Description string `form:"description"`
		ImageUrl    string `form:"image_url"`
	}

	var form RequestDataField

	if err := c.ShouldBind(&form); err != nil {
		fmt.Println(err)
		apiErr := errors.NewBadRequestError("Bad request")
		c.IndentedJSON(apiErr.Status, apiErr)
		return
	}

	// id,  := strconv.Atoi(c.Param("id"))
	id := form.ID
	word := form.Word
	description := form.Description
	imageUrl := form.ImageUrl

	fmt.Printf("Updating a keyword id: %d", id)
	keyword, err := ku.keywordUseCase.GetKeyword(id)
	if err != nil {
		fmt.Println(err)
		apiErr := errors.NotFoundError("Not found")
		c.IndentedJSON(apiErr.Status, apiErr)
		return
	}

	// word := c.Param("word")
	// description := c.Param("description")

	err = ku.keywordUseCase.UpdateKeyword(id, word, description, imageUrl)
	if err != nil {
		fmt.Println(err)
		apiErr := errors.InternalSeverError("Server Error")
		c.IndentedJSON(apiErr.Status, apiErr)
		return
	}

	c.IndentedJSON(http.StatusOK, keyword)
}

func (ku *keywordController) DeleteKeyword(c *gin.Context) {
	// 20221213 Deleted prior to implement immutable model
	// id, err := strconv.Atoi(c.Param("id"))
	// if err != nil {
	// 	fmt.Println(err)
	// 	apiErr := errors.NewBadRequestError("Bad request")
	// 	c.IndentedJSON(apiErr.Status, apiErr)
	// }

	type RequestDataField struct {
		ID string `form:"id" binding:"required"`
	}
	var form RequestDataField

	id := form.ID

	// keyword, err := ku.keywordUseCase.GetKeyword(id)
	// if err != nil {
	// 	fmt.Println(err)
	// 	apiErr := errors.NotFoundError("Not found")
	// 	c.IndentedJSON(apiErr.Status, apiErr)
	// }

	err := ku.keywordUseCase.DeleteKeyword(id)
	if err != nil {
		fmt.Println(err)
		apiErr := errors.InternalSeverError("Server Error due to immutable setting")
		c.IndentedJSON(apiErr.Status, apiErr)
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Deleted"})

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
