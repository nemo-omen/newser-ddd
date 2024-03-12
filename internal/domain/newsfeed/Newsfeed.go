package newsfeed

import (
	"encoding/json"
	"newser/internal/domain/article"
	"newser/internal/domain/entity"
	"newser/internal/domain/value"
)

type Newsfeed struct {
	id          value.ID
	title       value.Title
	siteUrl     value.Link
	feedUrl     value.Link
	description value.Description
	copyright   string
	language    string
	feedType    string
	image       entity.ImageId
	articles    []article.ArticleId
	authors     []entity.PersonId
	categories  []entity.CategoryId
	slug        string
}

type NewsfeedProps struct {
	Title       string
	SiteUrl     string
	FeedUrl     string
	Description string
	Copyright   string
	Language    string
	FeedType    string
	Image       entity.Image
}

func NewNewsfeed(props NewsfeedProps) (*Newsfeed, error) {
	slug, err := value.NewSlug(props.Title)
	if err != nil {
		return nil, err
	}

	newsfeed := &Newsfeed{
		id:          value.NewId(),
		title:       props.Title,
		siteUrl:     props.SiteUrl,
		feedUrl:     props.FeedUrl,
		description: props.Description,
		copyright:   props.Copyright,
		language:    props.Language,
		feedType:    props.FeedType,
		image:       props.Image.ID,
		articles:    []article.ArticleId{},
		authors:     []entity.PersonId{},
		categories:  []entity.CategoryId{},
		slug:        slug,
	}
	return newsfeed, nil
}

func (n *Newsfeed) JSON() []byte {
	type newsfeedJson struct {
		ID          value.ID            `json:"id"`
		Title       value.Title         `json:"title"`
		SiteUrl     value.Link          `json:"siteUrl"`
		FeedUrl     value.Link          `json:"feedUrl"`
		Description value.Description   `json:"description"`
		Copyright   string              `json:"copyRight"`
		Language    string              `json:"language"`
		FeedType    string              `json:"feedType"`
		Image       entity.ImageId      `json:"image"`
		Articles    []article.ArticleId `json:"articles"`
		Authors     []entity.PersonId   `json:"authors"`
		Categories  []entity.CategoryId `json:"categories"`
	}

	nfj := newsfeedJson{
		ID:          n.id,
		Title:       n.title,
		SiteUrl:     n.siteUrl,
		FeedUrl:     n.feedUrl,
		Description: n.description,
		Copyright:   n.copyright,
		Language:    n.language,
		FeedType:    n.feedType,
		Image:       n.image,
		Articles:    n.articles,
		Authors:     n.authors,
		Categories:  n.categories,
	}

	b, _ := json.MarshalIndent(nfj, "", "  ")
	return b
}

func (n *Newsfeed) String() string {
	return string(n.JSON())
}

func (n *Newsfeed) AddArticle(a article.ArticleId) {
	n.articles = append(n.articles, a)
}

func (n *Newsfeed) AddAuthor(p entity.PersonId) {
	n.authors = append(n.authors, p)
}

func (n *Newsfeed) AddCategory(c entity.CategoryId) {
	n.categories = append(n.categories, c)
}
