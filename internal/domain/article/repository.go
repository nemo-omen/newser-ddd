package article

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrNotFound   = errors.New("article not found")
	ErrNotSaved   = errors.New("article not saved")
	ErrNotUpdated = errors.New("article not updated")
	ErrNotDeleted = errors.New("article not deleted")
)

type ArticleRepository interface {
	Save(article *Article) error
	FindById(id uuid.UUID) (*Article, error)
	FindByTitle(title string) (*Article, error)
	FindBySlug(slug string) (*Article, error)
	FindByNewsfeedId(newsfeedId uuid.UUID) ([]*Article, error)
	Update(article *Article) error
	Delete(id string) error
}
