package annotation

import (
	"newser/internal/domain/value"
)

type Annotation struct {
	id      value.AnnotationId
	title   string
	content string
	article value.ArticleId
	// location
	// created
	// updated
	// collection
}
