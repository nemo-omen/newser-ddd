package newsfeed

import "errors"

var (
	ErrNotFound   = errors.New("newsfeed not found")
	ErrNotSaved   = errors.New("newsfeed not saved")
	ErrNotUpdated = errors.New("newsfeed not updated")
	ErrNotDeleted = errors.New("newsfeed not deleted")
)

type NewsfeedRepository interface {
	Save(newsfeed *Newsfeed) error
	FindById(id string) (*Newsfeed, error)
	FindAll() ([]*Newsfeed, error)
	FindByTitle(title string) (*Newsfeed, error)
	FindBySlug(slug string) (*Newsfeed, error)
	Delete(id string) error
}
