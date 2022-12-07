package repository

import (
	"profileyou/api/domain/model"
)

type KeywordRepository interface {
	GetKeyword(id int) model.Keyword
	GetKeywords() []model.Keyword
	Create(k model.Keyword)
	Update(k model.Keyword)
	Delete(k model.Keyword)
}
