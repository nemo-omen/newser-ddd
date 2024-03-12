package collection

import (
	"newser/internal/domain/value"
)

type Collection struct {
	id    value.ID
	title value.Title
	user  value.ID
	items []value.ID
	slug  value.Slug
}
