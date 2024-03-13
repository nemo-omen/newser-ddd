package collection

import (
	"encoding/json"
	"newser/internal/domain/value"
)

type Collection struct {
	id       value.CollectionId
	title    value.Title
	user     value.PersonId
	articles []value.ArticleId
	slug     value.Slug
}

type CollectionProps struct {
	Title string
	User  value.PersonId
}

func NewCollection(props CollectionProps) (*Collection, error) {
	validTitle, err := value.NewTitle(props.Title)
	if err != nil {
		return nil, err
	}

	slug, err := value.NewSlug(props.Title)
	if err != nil {
		return nil, err
	}

	return &Collection{
		id:       value.CollectionId{ID: value.NewId()},
		title:    validTitle,
		user:     props.User,
		articles: []value.ArticleId{},
		slug:     slug,
	}, nil
}

func (c *Collection) AddArticle(article value.ArticleId) {
	c.articles = append(c.articles, article)
}

func (c *Collection) RemoveArticle(article value.ArticleId) {
	for i, a := range c.articles {
		if a == article {
			c.articles = append(c.articles[:i], c.articles[i+1:]...)
		}
	}
}

func (c *Collection) ID() value.CollectionId {
	return c.id
}

func (c *Collection) Title() value.Title {
	return c.title
}

func (c *Collection) User() value.PersonId {
	return c.user
}

func (c *Collection) Articles() []value.ArticleId {
	return c.articles
}

func (c *Collection) Slug() value.Slug {
	return c.slug
}

func (c *Collection) SetTitle(title string) {
	title, err := value.NewTitle(title)
	if err != nil {
		return
	}
	slug, err := value.NewSlug(title)
	if err != nil {
		return
	}
	c.title = title
	c.slug = slug
}

func (c *Collection) JSON() []byte {
	type jsonCollection struct {
		ID       value.CollectionId `json:"id"`
		Title    value.Title        `json:"title"`
		User     value.PersonId     `json:"user"`
		Articles []value.ArticleId  `json:"articles"`
		Slug     value.Slug         `json:"slug"`
	}
	jc := jsonCollection{
		ID:       c.id,
		Title:    c.title,
		User:     c.user,
		Articles: c.articles,
		Slug:     c.slug,
	}
	b, _ := json.MarshalIndent(jc, "", "  ")
	return b
}

func (c *Collection) String() string {
	return string(c.JSON())
}
