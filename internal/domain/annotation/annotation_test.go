package annotation

import (
	"newser/internal/domain/value"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnnotation(t *testing.T) {
	artId := value.ArticleId{ID: value.NewId()}
	annotation, err := NewAnnotation("Test Annotation", "This is a test annotation", artId)

	if err != nil {
		t.Errorf("Error creating annotation: %v", err)
	}

	assert.Equal(t, "Test Annotation", annotation.title, "should have correct title")
	assert.Equal(t, "This is a test annotation", annotation.content, "should have correct content")
	assert.Equal(t, artId, annotation.article, "should have correct article")
	assert.NotEqual(t, value.AnnotationId{}, annotation.id, "should have a valid id")
}
