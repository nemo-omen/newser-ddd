package collection

import "errors"

var (
	ErrNotFound   = errors.New("collection not found")
	ErrNotSaved   = errors.New("collection not saved")
	ErrNotUpdated = errors.New("collection not updated")
	ErrNotDeleted = errors.New("collection not deleted")
)

type CollectionRepository interface {
	Save(collection *Collection) error
	FindById(id string) (*Collection, error)
	FindBySlug(slug string) (*Collection, error)
	FindAll(userId string) ([]*Collection, error)
	FindAllByUserId(userId string) ([]*Collection, error)
	Update(collection *Collection) error
	Delete(id string) error
}
