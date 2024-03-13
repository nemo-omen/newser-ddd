package newsfeed

import (
	"encoding/json"
	"newser/internal/domain/value"
)

type Newsfeed struct {
	id          value.NewsfeedId
	title       value.Title
	siteUrl     value.Link
	feedUrl     value.Link
	description value.Description
	copyright   string
	language    string
	feedType    string
	image       value.ImageId
	articles    []value.ArticleId
	authors     []value.PersonId
	categories  []value.CategoryId
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
	Image       value.ImageId
}

func NewNewsfeed(props NewsfeedProps) (*Newsfeed, error) {
	slug, err := value.NewSlug(props.Title)
	if err != nil {
		return nil, err
	}

	validTitle, err := value.NewTitle(props.Title)
	if err != nil {
		return nil, err
	}

	siteUrl, err := value.NewLink(props.SiteUrl)
	if err != nil {
		return nil, err
	}

	feedUrl, err := value.NewLink(props.FeedUrl)
	if err != nil {
		return nil, err
	}

	desc := value.NewDescription(props.Description)

	newsfeed := &Newsfeed{
		id:          value.NewsfeedId{ID: value.NewId()},
		title:       validTitle,
		siteUrl:     siteUrl,
		feedUrl:     feedUrl,
		description: desc,
		copyright:   props.Copyright,
		language:    props.Language,
		feedType:    props.FeedType,
		image:       props.Image,
		articles:    []value.ArticleId{},
		authors:     []value.PersonId{},
		categories:  []value.CategoryId{},
		slug:        slug,
	}
	return newsfeed, nil
}

func (n *Newsfeed) JSON() []byte {
	type newsfeedJson struct {
		ID          value.NewsfeedId   `json:"id"`
		Title       value.Title        `json:"title"`
		SiteUrl     value.Link         `json:"siteUrl"`
		FeedUrl     value.Link         `json:"feedUrl"`
		Description value.Description  `json:"description"`
		Copyright   string             `json:"copyRight"`
		Language    string             `json:"language"`
		FeedType    string             `json:"feedType"`
		Image       value.ImageId      `json:"image"`
		Articles    []value.ArticleId  `json:"articles"`
		Authors     []value.PersonId   `json:"authors"`
		Categories  []value.CategoryId `json:"categories"`
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

func (n *Newsfeed) ID() value.NewsfeedId {
	return n.id
}

func (n *Newsfeed) String() string {
	return string(n.JSON())
}

func (n *Newsfeed) AddArticle(a value.ArticleId) {
	n.articles = append(n.articles, a)
}

func (n *Newsfeed) AddAuthor(p value.PersonId) {
	n.authors = append(n.authors, p)
}

func (n *Newsfeed) AddCategory(c value.CategoryId) {
	n.categories = append(n.categories, c)
}
