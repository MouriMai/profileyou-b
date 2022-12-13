package keyword

import (
	"fmt"

	"github.com/google/uuid"
)

type Keyword struct {
	keywordId   keywordId
	word        word
	description description
	imageUrl    imageUrl
}

type keywordId string
type word string
type description string
type imageUrl string

func New(keywordId string, word string, description string, imageUrl string) (*Keyword, error) {
	createdKeywordId, err := NewKeywordId(keywordId)
	if err != nil {
		return nil, err
	}

	createdWord, err := newWord(word)
	if err != nil {
		return nil, err
	}

	createdDescription, err := newDescription(description)
	if err != nil {
		return nil, err
	}
	createdImageUrl, err := newImageUrl(imageUrl)
	if err != nil {
		return nil, err
	}

	keyword := Keyword{
		keywordId:   *createdKeywordId,
		word:        *createdWord,
		description: *createdDescription,
		imageUrl:    *createdImageUrl,
	}

	return &keyword, nil
}

// Create Keyword
func Create(word string, description string, imageUrl string) (*Keyword, error) {
	keywordId := uuid.New().String()
	keyword, err := New(keywordId, word, description, imageUrl)

	if err != nil {
		return nil, err
	}

	return keyword, err
}

// Getter
func (k Keyword) GetKeywordId() string {
	return string(k.keywordId)
}

func (k Keyword) GetWord() string {
	return string(k.word)
}

func (k Keyword) GetDescription() string {
	return string(k.description)
}

func (k Keyword) GetImageUrl() string {
	return string(k.imageUrl)
}

// value constructors
func NewKeywordId(value string) (*keywordId, error) {
	if value == "" {
		err := fmt.Errorf("%s", "empty arg:keywordId NewKeywordId()")
		return nil, err
	}

	keywordId := keywordId(value)

	return &keywordId, nil
}

func newWord(value string) (*word, error) {
	if value == "" {
		err := fmt.Errorf("%s", "empty arg:word newWord()")
		return nil, err
	}

	word := word(value)

	return &word, nil
}

func newDescription(value string) (*description, error) {
	if value == "" {
		err := fmt.Errorf("%s", "empty arg:description newDescription()")
		return nil, err
	}

	description := description(value)

	return &description, nil
}

func newImageUrl(value string) (*imageUrl, error) {
	if value == "" {
		err := fmt.Errorf("%s", "empty arg:imageUrl newImageUrl()")
		return nil, err
	}

	imageUrl := imageUrl(value)

	return &imageUrl, nil
}
