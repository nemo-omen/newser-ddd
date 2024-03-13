package collection

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrNotFound   = errors.New("collection not found")
	ErrNotSaved   = errors.New("collection not saved")
	ErrNotUpdated = errors.New("collection not updated")
	ErrNotDeleted = errors.New("collection not deleted")
)

type CollectionRepository interface {
	Save(collection *Collection) error
	FindById(id uuid.UUID) (*Collection, error)
	FindBySlug(slug string) (*Collection, error)
	FindAllByUserId(userId uuid.UUID) ([]*Collection, error)
	Update(collection *Collection) error
	Delete(id uuid.UUID) error
}
