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

func NewAnnotation(title string, content string, article value.ArticleId) (*Annotation, error) {
	validTitle, err := value.NewTitle(title)
	if err != nil {
		return nil, err
	}

	return &Annotation{
		id:      value.AnnotationId{ID: value.NewId()},
		title:   validTitle,
		content: content,
		article: article,
	}, nil
}
