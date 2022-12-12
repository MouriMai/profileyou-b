package repository

import (
	"profileyou/api/domain/model"
)

type KeywordRepository interface {
	GetKeyword(id int) (result *model.Keyword, err error)
	GetKeywords() (result []model.Keyword, err error)
	Create(k model.Keyword) error
	Update(k model.Keyword) error
	Delete(k model.Keyword) error
}
