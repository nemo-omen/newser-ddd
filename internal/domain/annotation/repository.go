package annotation

import "errors"

var (
	ErrNotFound   = errors.New("annotation not found")
	ErrNotSaved   = errors.New("annotation not saved")
	ErrNotUpdated = errors.New("annotation not updated")
	ErrNotDeleted = errors.New("annotation not deleted")
)

type AnnotationRepository interface {
	Save(annotation *Annotation) error
	FindById(id string) (*Annotation, error)
	FindByTitle(title string) (*Annotation, error)
	FindAll(userId string) ([]*Annotation, error)
	Update(annotation *Annotation) error
	Delete(id string) error
}
