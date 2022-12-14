package model

import (
	"gorm.io/gorm"
)

type Keyword struct {
	gorm.Model
	Word        string `json:"word`
	Description string `json:"description`
	ImageUrl    string `json:"image_url`
	KeywordId   string `json:"keyword_id`
}
