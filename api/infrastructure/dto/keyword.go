package dto

import (
	"profileyou/api/domain/model/keyword"
	"time"
)

type Keyword struct {
	ID          int
	KeywordId   string
	Word        string
	Description string
	ImageUrl    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

func ConvertKeyword(k *keyword.Keyword) *Keyword {
	return &Keyword{
		KeywordId:   k.GetKeywordId(),
		Word:        k.GetWord(),
		Description: k.GetDescription(),
		ImageUrl:    k.GetImageUrl(),
	}
}

func AdaptKeyword(converted_keyword *Keyword) (*keyword.Keyword, error) {
	keyword, err := keyword.New(
		converted_keyword.KeywordId,
		converted_keyword.Word,
		converted_keyword.Description,
		converted_keyword.ImageUrl,
	)

	if err != nil {
		return nil, err
	}

	return keyword, nil
}

func AdaptKeywords(converted_keywords []*Keyword) ([]*keyword.Keyword, error) {
	var keywords []*keyword.Keyword

	for _, converted_keyword := range converted_keywords {
		keyword, err := keyword.New(
			converted_keyword.KeywordId,
			converted_keyword.Word,
			converted_keyword.Description,
			converted_keyword.ImageUrl,
		)

		if err != nil {
			return nil, err
		}
		keywords = append(keywords, keyword)
	}

	return keywords, nil
}
