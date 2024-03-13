package article

import (
	"encoding/json"
	"newser/internal/domain/value"
	"time"

	"github.com/google/uuid"
)

type Article struct {
	id              value.ArticleId
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
	authors         []value.PersonId
	image           value.ImageId
	annotations     []value.AnnotationId
	categories      []value.CategoryId
	feed            value.NewsfeedId
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
	valFeedId, err := value.IdFromString(props.FeedId)
	if err != nil {
		return nil, err
	}
	validTitle, err := value.NewTitle(props.Title)
	if err != nil {
		return nil, err
	}
	// sanitize content & add 'target="_blank"' to links
	preparedContent := value.NewItemContent(props.Content)

	desc := ""
	// if item content is long enough
	// set it as description
	if len(preparedContent) > 240 {
		desc = value.NewDescription(preparedContent[:240] + "...")
	} else {
		// otherwise, just use default description
		desc = value.NewDescription(props.Description)
	}

	l, _ := value.NewLink(props.Link)
	ll := []value.Link{}

	for _, href := range props.Links {
		lnk, _ := value.NewLink(href)
		if len(lnk) > 0 {
			ll = append(ll, lnk)
		}
	}

	var valImgId value.ID
	if props.Image != "" && props.Image != uuid.Nil.String() {
		valImgId, err = value.IdFromString(props.Image)
		if err != nil {
			return nil, err
		}
	}

	slug, err := value.NewSlug(props.Title)
	if err != nil {
		return nil, err
	}

	article := &Article{
		id:              value.ArticleId{ID: value.NewId()},
		title:           validTitle,
		description:     desc,
		content:         preparedContent,
		link:            l,
		links:           ll,
		updated:         props.Updated,
		updatedParsed:   props.UpdatedParsed,
		published:       props.Published,
		publishedParsed: props.PublishedParsed,
		guid:            props.GUID,
		authors:         []value.PersonId{},
		image:           value.ImageId{ID: valImgId},
		annotations:     []value.AnnotationId{},
		categories:      []value.CategoryId{},
		feed:            value.NewsfeedId{ID: valFeedId},
		slug:            slug,
	}

	return article, nil
}

func (a *Article) JSON() []byte {
	type jsonArticle struct {
		ID              value.ArticleId   `json:"id"`
		Title           value.Title       `json:"title"`
		Description     value.Description `json:"description"`
		Content         value.ItemContent `json:"content"`
		Link            value.Link        `json:"link"`
		Links           []value.Link      `json:"links"`
		Updated         string            `json:"updated"`
		UpdatedParsed   time.Time         `json:"updatedParsed"`
		Published       string            `json:"published"`
		PublishedParsed time.Time         `json:"publishedParsed"`
		GUID            string            `json:"guid"`
		Image           value.ImageId     `json:"image"`
		FeedId          value.NewsfeedId  `json:"feedId"`
		Slug            value.Slug        `json:"slug"`
		Read            bool              `json:"read"`
		Saved           bool              `json:"saved"`
	}

	ja := jsonArticle{
		ID:              a.id,
		Title:           a.title,
		Content:         a.content,
		Description:     a.description,
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

func (a *Article) ID() value.ArticleId {
	return a.id
}
