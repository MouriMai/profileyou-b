package repository

import (
	"profileyou/api/domain/model/keyword"
)

type KeywordRepository interface {
	GetKeyword(id string) (result *keyword.Keyword, err error)
	GetKeywords() (result []*keyword.Keyword, err error)
	Create(k *keyword.Keyword) error
	Update(k *keyword.Keyword) error
	Delete(k *keyword.Keyword) error
}
