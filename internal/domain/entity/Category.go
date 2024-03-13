package entity

import (
	"encoding/json"
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

func (c *Category) JSON() []byte {
	type jsonCategory struct {
		ID   value.CategoryId `json:"id"`
		Term string           `json:"term"`
	}
	jc := jsonCategory{
		ID:   c.ID,
		Term: c.Term,
	}
	b, _ := json.MarshalIndent(jc, "", "  ")
	return b
}

func (c *Category) String() string {
	return string(c.JSON())
}
