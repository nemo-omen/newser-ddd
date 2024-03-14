package category

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrNotFound   = errors.New("category not found")
	ErrNotSaved   = errors.New("category not saved")
	ErrNotUpdated = errors.New("category not updated")
	ErrNotDeleted = errors.New("category not deleted")
)

type CategoryRepository interface {
	Save(category *Category) error
	FindById(id uuid.UUID) (*Category, error)
	FindByTerm(term string) (*Category, error)
	FindAll() ([]*Category, error)
}
