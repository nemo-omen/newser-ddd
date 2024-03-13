package newsfeed

import (
	"newser/internal/domain/value"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNewsfeed(t *testing.T) {
	props := NewsfeedProps{
		Title:       "Test Newsfeed",
		SiteUrl:     "https://example.com",
		FeedUrl:     "https://example.com/feed",
		Description: "Test description",
		Copyright:   "© 2022",
		Language:    "en",
		FeedType:    "rss",
		Image:       value.ImageId{ID: value.NewId()},
	}

	newsfeed, err := NewNewsfeed(props)

	assert.NoError(t, err)
	assert.NotNil(t, newsfeed)
	assert.NotEmpty(t, newsfeed.id)
	assert.Equal(t, props.Title, newsfeed.title)
	assert.Equal(t, props.SiteUrl, newsfeed.siteUrl)
	assert.Equal(t, props.FeedUrl, newsfeed.feedUrl)
	assert.Equal(t, props.Description, newsfeed.description)
	assert.Equal(t, props.Copyright, newsfeed.copyright)
	assert.Equal(t, props.Language, newsfeed.language)
	assert.Equal(t, props.FeedType, newsfeed.feedType)
	assert.Equal(t, props.Image, newsfeed.image)
	assert.Equal(t, "test-newsfeed", newsfeed.slug)
	assert.Empty(t, newsfeed.articles)
	assert.Empty(t, newsfeed.authors)
	assert.Empty(t, newsfeed.categories)
}

func TestNewsfeedAddArticle(t *testing.T) {
	props := NewsfeedProps{
		Title:       "Test Newsfeed",
		SiteUrl:     "https://example.com",
		FeedUrl:     "https://example.com/feed",
		Description: "Test description",
		Copyright:   "© 2022",
		Language:    "en",
		FeedType:    "rss",
		Image:       value.ImageId{ID: value.NewId()},
	}

	aid1 := value.ArticleId{ID: value.NewId()}
	aid2 := value.ArticleId{ID: value.NewId()}
	aid3 := value.ArticleId{ID: value.NewId()}

	newsfeed, err := NewNewsfeed(props)
	assert.NoError(t, err)

	newsfeed.AddArticle(aid1)
	assert.Len(t, newsfeed.articles, 1)
	newsfeed.AddArticle(aid2)
	assert.Len(t, newsfeed.articles, 2)
	newsfeed.AddArticle(aid3)
	assert.Len(t, newsfeed.articles, 3)
}
