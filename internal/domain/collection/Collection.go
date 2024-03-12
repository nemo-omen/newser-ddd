package collection

import (
	"newser/internal/domain/value"
)

type Collection struct {
	id       value.CollectionId
	title    value.Title
	user     value.PersonId
	articles []value.ArticleId
	slug     value.Slug
}
