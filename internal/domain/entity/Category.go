package entity

import (
	"newser/internal/domain/value"
)

type Category struct {
	ID   value.CategoryId
	Term string
}

func NewCategory(term string) (*Category, error) {
	validTerm, err := value.NewTerm(term)
	if err != nil {
		return nil, err
	}

	category := &Category{
		ID:   value.CategoryId{ID: value.NewId()},
		Term: validTerm,
	}
	return category, nil
}
