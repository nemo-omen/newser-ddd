package article

import (
	"encoding/json"
	"newser/internal/domain/value"
	"time"
)

type Article struct {
	id              value.ID
	title           value.Title
	description     value.Description
	content         value.ItemContent
	link            value.Link
	links           []value.Link
	updated         string
	updatedParsed   time.Time
	published       string
	publishedParsed time.Time
	guid            string
	authors         []value.ID
	image           value.ID
	annotations     []value.ID
	categories      []value.ID
	feed            value.ID
	slug            string
	read            bool
	saved           bool
}

type ArticleProps struct {
	Title           string
	Description     string
	Content         string
	Link            string
	Links           []string
	Updated         string
	UpdatedParsed   time.Time
	Published       string
	PublishedParsed time.Time
	GUID            string
	Image           string
	FeedId          string
}

func NewArticle(props ArticleProps) (*Article, error) {
	slug, err := value.NewSlug(props.Title)
	if err != nil {
		return nil, err
	}

	article := &Article{
		title:           props.Title,
		description:     props.Description,
		content:         props.Content,
		link:            props.Link,
		links:           props.Links,
		updated:         props.Updated,
		updatedParsed:   props.UpdatedParsed,
		published:       props.Published,
		publishedParsed: props.PublishedParsed,
		guid:            props.GUID,
		authors:         []value.ID{},
		image:           value.IdFromString(props.Image),
		annotations:     []value.ID{},
		categories:      []value.ID{},
		feed:            value.IdFromString(props.FeedId),
		slug:            slug,
	}

	return article, nil
}

func (a *Article) JSON() []byte {
	type jsonArticle struct {
		ID              value.ID          `json:"id"`
		Title           value.Title       `json:"title"`
		Content         value.ItemContent `json:"content"`
		Link            value.Link        `json:"link"`
		Links           []value.Link      `json:"links"`
		Updated         string            `json:"updated"`
		UpdatedParsed   time.Time         `json:"updatedParsed"`
		Published       string            `json:"published"`
		PublishedParsed time.Time         `json:"publishedParsed"`
		GUID            string            `json:"guid"`
		Image           value.ID          `json:"image"`
		FeedId          value.ID          `json:"feedId"`
		Slug            value.Slug        `json:"slug"`
		Read            bool              `json:"read"`
		Saved           bool              `json:"saved"`
	}

	ja := jsonArticle{
		ID:              a.id,
		Title:           a.title,
		Content:         a.content,
		Link:            a.link,
		Updated:         a.updated,
		UpdatedParsed:   a.updatedParsed,
		Published:       a.published,
		PublishedParsed: a.publishedParsed,
		GUID:            a.guid,
		Image:           a.image,
		FeedId:          a.feed,
		Slug:            a.slug,
		Read:            a.read,
		Saved:           a.saved,
	}

	j, _ := json.MarshalIndent(ja, "", "  ")
	return j
}

func (a *Article) String() string {
	return string(a.JSON())
}

func (a *Article) Read() {
	a.read = true
}

func (a *Article) Unread() {
	a.read = false
}

func (a *Article) Save() {
	a.saved = true
}

func (a *Article) Unsave() {
	a.saved = false
}
