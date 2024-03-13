package collection

import (
	"newser/internal/domain/value"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testSlug, _ = value.NewSlug("Test Collection")

func TestCollection(t *testing.T) {
	c := Collection{
		id:       value.CollectionId{ID: value.NewId()},
		title:    value.Title("Test Collection"),
		user:     value.PersonId{ID: value.NewId()},
		articles: []value.ArticleId{{ID: value.NewId()}, {ID: value.NewId()}},
		slug:     testSlug,
	}

	assert.Equal(t, "Test Collection", string(c.title), "should have correct title")
	assert.NotEqual(t, value.CollectionId{}, c.id, "should have a valid id")
	assert.NotEqual(t, value.PersonId{}, c.user, "should have a valid user")
	assert.NotEqual(t, 0, len(c.articles), "should have articles")
	assert.Equal(t, testSlug, c.slug, "should have correct slug")
	assert.Equal(t, 2, len(c.articles), "should have 2 articles")
}

func TestAddArticle(t *testing.T) {
	c := Collection{
		id:       value.CollectionId{ID: value.NewId()},
		title:    value.Title("Test Collection"),
		user:     value.PersonId{ID: value.NewId()},
		articles: []value.ArticleId{{ID: value.NewId()}, {ID: value.NewId()}},
		slug:     testSlug,
	}

	c.AddArticle(value.ArticleId{ID: value.NewId()})

	assert.Equal(t, 3, len(c.articles), "should have 3 articles")
}

func TestRemoveArticle(t *testing.T) {
	c := Collection{
		id:       value.CollectionId{ID: value.NewId()},
		title:    value.Title("Test Collection"),
		user:     value.PersonId{ID: value.NewId()},
		articles: []value.ArticleId{{ID: value.NewId()}, {ID: value.NewId()}},
		slug:     testSlug,
	}

	c.RemoveArticle(c.articles[0])

	assert.Equal(t, 1, len(c.articles), "should have 3 articles")
}
