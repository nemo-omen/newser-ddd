package article

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var shortProps = ArticleProps{
	FeedId:          uuid.New().String(),
	Title:           "Test Article",
	Description:     "This is a test article",
	Content:         "<p>Lorem ipsum dolor sit amet</p>",
	Link:            "https://example.com",
	Links:           []string{"https://example.com/1", "https://example.com/2"},
	Updated:         "2022-01-01T00:00:00Z",
	UpdatedParsed:   time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
	Published:       "2022-01-01T00:00:00Z",
	PublishedParsed: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
	GUID:            "123456789",
	Image:           uuid.New().String(),
}

var longProps = ArticleProps{
	FeedId:          uuid.New().String(),
	Title:           "Test Article",
	Description:     "This is a test article",
	Content:         "<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>",
	Link:            "https://example.com",
	Links:           []string{"https://example.com/1", "https://example.com/2"},
	Updated:         "2022-01-01T00:00:00Z",
	UpdatedParsed:   time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
	Published:       "2022-01-01T00:00:00Z",
	PublishedParsed: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
	GUID:            "123456789",
	Image:           uuid.New().String(),
}

var longDescription = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis ..."

func TestShortArticle(t *testing.T) {

	article, err := NewArticle(shortProps)

	assert.NoError(t, err)
	assert.NotNil(t, article)
	assert.NotEmpty(t, article.ID)
	assert.Equal(t, shortProps.Title, article.title)
	assert.Equal(t, shortProps.Description, article.description)
	assert.Equal(t, shortProps.Content, article.content)
	assert.Equal(t, shortProps.Link, article.link)
	assert.Len(t, article.links, len(shortProps.Links))
	assert.Equal(t, shortProps.Updated, article.updated)
	assert.Equal(t, shortProps.UpdatedParsed, article.updatedParsed)
	assert.Equal(t, shortProps.Published, article.published)
	assert.Equal(t, shortProps.PublishedParsed, article.publishedParsed)
	assert.Equal(t, shortProps.GUID, article.guid)
	assert.Empty(t, article.authors)
	assert.Equal(t, shortProps.Image, article.image.ID.String())
	assert.Empty(t, article.annotations)
	assert.Empty(t, article.categories)
	assert.Equal(t, shortProps.FeedId, article.feed.ID.String())
	assert.Equal(t, "test-article", article.slug, "should have correct slug")
}

func TestLongArticle(t *testing.T) {

	article, err := NewArticle(longProps)

	assert.NoError(t, err)
	assert.NotNil(t, article)
	assert.NotEmpty(t, article.ID)
	assert.Equal(t, longProps.Title, article.title)
	assert.Equal(t, longDescription, article.description)
	assert.Equal(t, longProps.Content, article.content)
	assert.Equal(t, longProps.Link, article.link)
	assert.Len(t, article.links, len(longProps.Links))
	assert.Equal(t, longProps.Updated, article.updated)
	assert.Equal(t, longProps.UpdatedParsed, article.updatedParsed)
	assert.Equal(t, longProps.Published, article.published)
	assert.Equal(t, longProps.PublishedParsed, article.publishedParsed)
	assert.Equal(t, longProps.GUID, article.guid)
	assert.Empty(t, article.authors)
	assert.Equal(t, longProps.Image, article.image.ID.String())
	assert.Empty(t, article.annotations)
	assert.Empty(t, article.categories)
	assert.Equal(t, longProps.FeedId, article.feed.ID.String())
	assert.Equal(t, "test-article", article.slug, "should have correct slug")
}
