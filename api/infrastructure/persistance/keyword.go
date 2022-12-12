package persistance

import (
	"profileyou/api/domain/model"
	"profileyou/api/domain/repository"

	"gorm.io/gorm"
)

type keywordPersistance struct {
	Conn *gorm.DB
}

func NewKeywordPersistance(conn *gorm.DB, k repository.KeywordRepository) *keywordPersistance {
	return &keywordPersistance{Conn: conn}
}

func (kr *keywordPersistance) GetKeyword(id int) (result *model.Keyword, err error) {

	var keyword model.Keyword
	if result := kr.Conn.First(&keyword, id); result.Error != nil {
		err := result.Error
		return nil, err
	}

	return &keyword, nil

}

func (kr *keywordPersistance) GetKeywords() (result []model.Keyword, err error) {

	var keywords []model.Keyword

	if result := kr.Conn.Find(&keywords); result.Error != nil {
		err := result.Error
		return nil, err
	}

	return keywords, nil
}

func (kr *keywordPersistance) Create(k model.Keyword) error {

	kr.Conn.Create(&k)
	return nil
}

func (kr *keywordPersistance) Update(k model.Keyword) error {

	kr.Conn.Save(&k)
	return nil

}

func (kr *keywordPersistance) Delete(k model.Keyword) error {

	kr.Conn.Delete(&k)
	return nil
}
