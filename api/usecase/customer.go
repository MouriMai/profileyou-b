package usecase

import (
	"profileyou/api/domain/model"
	"profileyou/api/domain/repository"
)

type KeywordUseCase interface {
	// GetKeyword(id int) (result *model.Keyword, err error)
	// GetKeywords() (result []model.Keyword, err error)
	// Create(k model.Keyword) error
	// Update(k model.Keyword) error
	// Delete(k model.Keyword) error
	GetKeyword(id int) (result *model.Keyword, err error)
	GetKeywords() (result []model.Keyword, err error)
	CreateKeyword(word string) error
	UpdateKeyword(id int, word string, image_url string) error
	DeleteKeyword(id int) error
}

type keywordUseCase struct {
	keywordRepository repository.KeywordRepository
}

func NewKeywordUseCase(kr repository.KeywordRepository) KeywordUseCase {
	return &keywordUseCase{
		keywordRepository: kr,
	}
}

func (ku *keywordUseCase) GetKeyword(id int) (result *model.Keyword, err error) {
	keyword, err := ku.keywordRepository.GetKeyword(id)
	if err != nil {
		return nil, err
	}

	return keyword, nil
}

func (ku *keywordUseCase) GetKeywords() (result []model.Keyword, err error) {
	keywords, err := ku.keywordRepository.GetKeywords()
	if err != nil {
		return nil, err
	}

	return keywords, nil
}

func (ku *keywordUseCase) CreateKeyword(word string) error {
	keyword := model.Keyword{Word: word}
	err := ku.keywordRepository.Create(keyword)
	if err != nil {
		return err
	}

	return nil
}

func (ku *keywordUseCase) UpdateKeyword(id int, word string, description string) error {
	keyword, err := ku.keywordRepository.GetKeyword(id)
	if err != nil {
		return err
	}

	keyword.Word = word
	keyword.Description = description
	err = ku.keywordRepository.Update(*keyword)
	if err != nil {
		return err
	}

	return nil
}

func (ku *keywordUseCase) DeleteKeyword(id int) error {
	keyword, err := ku.keywordRepository.GetKeyword(id)
	if err != nil {
		return err
	}

	err = ku.keywordRepository.Delete(*keyword)
	if err != nil {
		return err
	}

	return nil
}
