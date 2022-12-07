package persistance

import (
	"profileyou/api/domain/model"
	"profileyou/api/domain/repository"

	"gorm.io/gorm"
)

type keywordPersistance struct {
	Conn *gorm.DB
}

func NewCustomerPersistance(conn *gorm.DB) repository.KeywordRepository {
	return &keywordPersistance{Conn: conn}
}

func (kr *keywordPersistance) GetKeyword(id int) model.Keyword {

	var keyword model.Keyword
	kr.Conn.First(&keyword, id)

	return keyword

}

func (kr *keywordPersistance) GetKeywords() []model.Keyword {

	var keywords []model.Keyword
	kr.Conn.Find(&keywords)

	return keywords
}

func (kr *keywordPersistance) Create(k model.Keyword) {

	kr.Conn.Create(&k)

}

func (kr *keywordPersistance) Update(k model.Keyword) {

	kr.Conn.Save(&k)

}

func (kr *keywordPersistance) Delete(k model.Keyword) {

	kr.Conn.Delete(&k)

}
