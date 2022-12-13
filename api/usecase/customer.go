package usecase

import (
	"fmt"
	"profileyou/api/domain/model/keyword"
	"profileyou/api/domain/repository"
)

type KeywordUseCase interface {
	// 221212
	// GetKeyword(id int) (result *model.Keyword, err error)
	// GetKeywords() (result []model.Keyword, err error)
	// Create(k model.Keyword) error
	// Update(k model.Keyword) error
	// Delete(k model.Keyword) error
	// 221213
	// GetKeyword(id int) (result *model.Keyword, err error)
	// GetKeywords() (result []model.Keyword, err error)
	// DeleteKeyword(id int) error
	// UpdateKeyword(id int, word string, image_url string) error
	GetKeyword(id string) (result *keyword.Keyword, err error)
	GetKeywords() (result []*keyword.Keyword, err error)
	CreateKeyword(word string, description string, imgaeUrl string) error
	UpdateKeyword(id string, word string, description string, image_url string) error
	DeleteKeyword(id string) error
}

type keywordUseCase struct {
	keywordRepository repository.KeywordRepository
}

func NewKeywordUseCase(kr repository.KeywordRepository) KeywordUseCase {
	return &keywordUseCase{
		keywordRepository: kr,
	}
}

func (ku *keywordUseCase) GetKeyword(id string) (result *keyword.Keyword, err error) {
	keyword, err := ku.keywordRepository.GetKeyword(id)
	if err != nil {
		return nil, err
	}

	return keyword, nil
}

func (ku *keywordUseCase) GetKeywords() (result []*keyword.Keyword, err error) {
	keywords, err := ku.keywordRepository.GetKeywords()
	fmt.Printf("USECASE: RETRIEVE KEYWORDS, %v\n", keywords)
	if err != nil {
		return nil, err
	}

	return keywords, nil
}

func (ku *keywordUseCase) CreateKeyword(word string, description string, imageUrl string) error {
	// keyword := model.Keyword{Word: word}
	// err := ku.keywordRepository.Create(keyword)
	keyword, err := keyword.Create(word, description, imageUrl)
	if err != nil {
		return err
	}

	err = ku.keywordRepository.Create(keyword)
	if err != nil {
		return err
	}

	return nil
}

func (ku *keywordUseCase) UpdateKeyword(id string, word string, description string, imgeUrl string) error {
	current_keyword, err := ku.keywordRepository.GetKeyword(id)
	if err != nil {
		return err
	}

	// keyword.Word = word
	// keyword.Description = description
	// err = ku.keywordRepository.Update(*keyword)
	keywordId := string(current_keyword.GetKeywordId())
	update_keyword, err := keyword.New(keywordId, word, description, imgeUrl)
	if err != nil {
		return err
	}
	err = ku.keywordRepository.Update(update_keyword)
	if err != nil {
		return err
	}

	return nil
}

func (ku *keywordUseCase) DeleteKeyword(id string) error {
	keyword, err := ku.keywordRepository.GetKeyword(id)
	if err != nil {
		return err
	}

	err = ku.keywordRepository.Delete(keyword)
	if err != nil {
		return err
	}

	return nil
}
