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

func (ku *keywordController) GetAllKeywordsGin(c *gin.Context) {
	keywords, err := ku.keywordUseCase.GetKeywords()
	fmt.Printf("keywords :%v\n", keywords)
	if err != nil {
		fmt.Println(err)
		apiErr := errors.NewBadRequestError("Get all Bad Request")
		c.IndentedJSON(apiErr.Status, apiErr)
		return
	}
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
	// c.HTML(200, "index.html", gin.H{"keywords": data})
	c.IndentedJSON(http.StatusOK, data)
}

func (ku *keywordController) GetKeyword(c *gin.Context) {
	id := c.Param("id")
	keyword, err := ku.keywordUseCase.GetKeyword(id)
	fmt.Printf("keyword id%v\n", keyword)
	if err != nil {
		fmt.Printf("GET KEYWORD: Could not find id%v\n", err)
		apiErr := errors.NotFoundError("Tried to find the record but Not found")
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

func (ku *keywordController) Index(c *gin.Context) {
	keywords, err := ku.keywordUseCase.GetKeywords()
	fmt.Printf("keywords :%v\n", keywords)
	if err != nil {
		fmt.Println(err)
		c.HTML(500, "500.html", gin.H{"error": err.Error()})
		// apiErr := errors.NewBadRequestError("Index Bad Request")
		// c.IndentedJSON(apiErr.Status, apiErr)
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
	c.HTML(200, "index.html", gin.H{"keywords": data})
	// c.IndentedJSON(http.StatusOK, data)
}

func (ku *keywordController) DetailKeyword(c *gin.Context) {
	id := c.Param("id")
	keyword, err := ku.keywordUseCase.GetKeyword(id)
	if err != nil {
		fmt.Printf("Could not find id%v\n", err)
		c.HTML(500, "500.html", gin.H{"error": err.Error()})
		// apiErr := errors.NotFoundError("Not found")
		// c.IndentedJSON(apiErr.Status, apiErr)
		return
	}
	// c.IndentedJSON(http.StatusOK, keyword)

	type ResultDataField struct {
		KeywordId   string
		Word        string
		Description string
		ImageUrl    string
	}

	data := ResultDataField{
		KeywordId:   string(keyword.GetKeywordId()),
		Word:        string(keyword.GetWord()),
		Description: string(keyword.GetDescription()),
		ImageUrl:    string(keyword.GetImageUrl()),
	}
	c.HTML(200, "detail.html", gin.H{"keyword": data})
}

func (ku *keywordController) CreateKeyword(c *gin.Context) {
	// 221214 form bindingができなくて一旦Paramから抜き取る変な感じ
	// word := c.Param("word")
	// fmt.Printf("Paramで受け取れるか%v\n", word)

	// 20221213 - Validation for create
	type RequestDataField struct {
		Word        string `form:"word" binding:"required"`
		Description string `form:"description"`
		ImageUrl    string `form:"image_url"`
	}

	var form RequestDataField
	fmt.Printf("&jsonには何が入っているのか%v\n", &form)
	if err := c.ShouldBind(&form); err != nil {
		fmt.Printf("Error: %v\n", err)
		c.HTML(400, "400.html", gin.H{"error": err.Error()})
		// apiErr := errors.NewBadRequestError("Bad request on binding json")
		// c.IndentedJSON(apiErr.Status, apiErr)
		return
	}

	word := form.Word
	description := ""
	imageUrl := ""
	fmt.Printf("Receive a post: %s\n", word)

	// keyword := model.Keyword{Word: word}
	err := ku.keywordUseCase.CreateKeyword(word, description, imageUrl)
	if err != nil {
		fmt.Println(err)
		c.HTML(500, "500.html", gin.H{"error": err.Error()})
		// apiErr := errors.InternalSeverError("Server Error when posting")
		// c.IndentedJSON(apiErr.Status, apiErr)
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
		c.HTML(400, "400.html", gin.H{"error": err.Error()})
		// apiErr := errors.NewBadRequestError("Bad request")
		// c.IndentedJSON(apiErr.Status, apiErr)
		return
	}

	id := form.ID
	word := form.Word
	description := form.Description
	imageUrl := form.ImageUrl

	fmt.Printf("Updating a keyword id: %v", id)
	fmt.Printf("Updating a image_url: %v", imageUrl)

	err := ku.keywordUseCase.UpdateKeyword(id, word, description, imageUrl)
	if err != nil {
		fmt.Println(err)
		c.HTML(500, "500.html", gin.H{"error": err.Error()})
		// apiErr := errors.InternalSeverError("Server Error")
		// c.IndentedJSON(apiErr.Status, apiErr)
		return
	}

	// c.IndentedJSON(http.StatusOK, keyword)
	c.Redirect(301, "/")
}

func (ku *keywordController) DeleteKeyword(c *gin.Context) {
	type RequestDataField struct {
		ID string `form:"id" binding:"required"`
	}
	var form RequestDataField

	if err := c.ShouldBind(&form); err != nil {
		fmt.Println(err)
		c.HTML(400, "400.html", gin.H{"error": err.Error()})
		return
	}

	id := form.ID

	err := ku.keywordUseCase.DeleteKeyword(id)
	if err != nil {
		fmt.Println(err)
		c.HTML(500, "500.html", gin.H{"error": err.Error()})
		// apiErr := errors.InternalSeverError("Server Error")
		// c.IndentedJSON(apiErr.Status, apiErr)
		return
	}

	c.Redirect(301, "/")
}
