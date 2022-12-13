package persistance

import (
	"fmt"
	"profileyou/api/domain/model/keyword"
	"profileyou/api/domain/repository"
	"profileyou/api/infrastructure/dto"

	"gorm.io/gorm"
)

type keywordPersistance struct {
	Conn *gorm.DB
}

//	func NewKeywordPersistance(conn *gorm.DB, k repository.KeywordRepository) *keywordPersistance {
//		return &keywordPersistance{Conn: conn}
//	}
func NewKeywordPersistance(conn *gorm.DB) repository.KeywordRepository {
	return &keywordPersistance{Conn: conn}
}

func (kp *keywordPersistance) GetKeyword(id string) (result *keyword.Keyword, err error) {

	var keyword dto.Keyword
	if result := kp.Conn.Where("keyword_id = ?", id).First(&keyword, id); result.Error != nil {
		err := result.Error
		return nil, err
	}

	// return &keyword, nil
	result_keyword, err := dto.AdaptKeyword(&keyword)
	if err != nil {
		return nil, err
	}

	return result_keyword, nil

}

func (kp *keywordPersistance) GetKeywords() (result []*keyword.Keyword, err error) {

	var keywords []*dto.Keyword

	fmt.Printf("&keywords, %v\n", &keywords)
	tmp := kp.Conn.Find(&keywords)
	fmt.Println(tmp.Error)
	if result := kp.Conn.Find(&keywords); result.Error != nil {
		fmt.Printf("PERSISTANCE DB ALL: RETRIEVE KEYWORDS, %v\n", result)
		err := result.Error
		return nil, err
	}
	fmt.Printf("PERSISTANCE: RETRIEVE KEYWORDS, %v\n", result)

	result_keywords, err := dto.AdaptKeywords(keywords)
	if err != nil {
		return nil, err
	}

	return result_keywords, nil
}

func (kp *keywordPersistance) Create(k *keyword.Keyword) error {

	converted_keyword := dto.ConvertKeyword(k)
	if result := kp.Conn.Create(converted_keyword); result.Error != nil {
		err := result.Error
		return err
	}
	return nil
}

func (kp *keywordPersistance) Update(k *keyword.Keyword) error {

	converted_keyword := dto.ConvertKeyword(k)
	if result := kp.Conn.Save(converted_keyword); result.Error != nil {
		err := result.Error
		return err
	}
	return nil
}

func (kp *keywordPersistance) Delete(k *keyword.Keyword) error {

	converted_keyword := dto.ConvertKeyword(k)
	if result := kp.Conn.Where("keyword_id = ?", converted_keyword.KeywordId).Delete(converted_keyword); result.Error != nil {
		err := result.Error
		return err
	}
	return nil
}
