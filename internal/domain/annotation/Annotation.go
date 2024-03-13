package annotation

import (
	"encoding/json"
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

func (a *Annotation) JSON() []byte {
	type jsonAnnotation struct {
		ID      value.AnnotationId `json:"id"`
		Title   string             `json:"title"`
		Content string             `json:"content"`
		Article value.ArticleId    `json:"article"`
	}
	js := jsonAnnotation{
		ID:      a.id,
		Title:   a.title,
		Content: a.content,
		Article: a.article,
	}
	b, _ := json.MarshalIndent(js, "", "  ")
	return b
}

func (a *Annotation) String() string {
	return string(a.JSON())
}

func (a *Annotation) Title() string {
	return a.title
}

func (a *Annotation) Content() string {
	return a.content
}

func (a *Annotation) Article() value.ArticleId {
	return a.article
}

func (a *Annotation) ID() value.AnnotationId {
	return a.id
}

func (a *Annotation) SetTitle(title string) error {
	validTitle, err := value.NewTitle(title)
	if err != nil {
		return err
	}
	a.title = validTitle
	return nil
}

func (a *Annotation) SetContent(content string) {
	a.content = content
}
