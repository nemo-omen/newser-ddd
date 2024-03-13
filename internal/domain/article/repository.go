package article

import "errors"

var (
	ErrNotFound   = errors.New("article not found")
	ErrNotSaved   = errors.New("article not saved")
	ErrNotUpdated = errors.New("article not updated")
	ErrNotDeleted = errors.New("article not deleted")
)

type ArticleRepository interface {
	Save(article *Article) error
	FindById(id string) (*Article, error)
	FindByTitle(title string) (*Article, error)
	FindAll() ([]*Article, error)
	Update(article *Article) error
	Delete(id string) error
}
