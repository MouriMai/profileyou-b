package repository

import (
	"profileyou/api/domain/model/keyword"
)

type KeywordRepository interface {
	// GetKeyword(id int) (result *model.Keyword, err error)
	// GetKeywords() (result []model.Keyword, err error)
	// Create(k model.Keyword) error
	// Update(k model.Keyword) error
	// Delete(k model.Keyword) error
	GetKeyword(id string) (result *keyword.Keyword, err error)
	GetKeywords() (result []keyword.Keyword, err error)
	Create(k *keyword.Keyword) error
	Update(k *keyword.Keyword) error
	Delete(k *keyword.Keyword) error
}
