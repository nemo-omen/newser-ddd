package memory

import (
	"newser/internal/domain/annotation"

	"github.com/google/uuid"
)

type AnnotationMemoryRepo struct {
	annotations map[uuid.UUID]*annotation.Annotation
}

func NewAnnotationMemoryRepo() *AnnotationMemoryRepo {
	return &AnnotationMemoryRepo{
		annotations: make(map[uuid.UUID]*annotation.Annotation),
	}
}

func (r *AnnotationMemoryRepo) Save(a *annotation.Annotation) error {
	r.annotations[a.ID().ID] = a
	return nil
}

func (r *AnnotationMemoryRepo) FindById(id string) (*annotation.Annotation, error) {
	aid, err := uuid.Parse(id)
	if err != nil {
		return nil, annotation.ErrNotFound
	}

	a, ok := r.annotations[aid]
	if !ok {
		return nil, annotation.ErrNotFound
	}
	return a, nil
}

func (r *AnnotationMemoryRepo) FindByTitle(title string) (*annotation.Annotation, error) {
	for _, a := range r.annotations {
		if a.Title() == title {
			return a, nil
		}
	}
	return nil, annotation.ErrNotFound
}

func (r *AnnotationMemoryRepo) FindAll(userId string) ([]*annotation.Annotation, error) {
	annotations := make([]*annotation.Annotation, 0, len(r.annotations))
	for _, a := range r.annotations {
		if a.User().String() == userId {
			annotations = append(annotations, a)
		}
	}
	return annotations, nil
}

func (r *AnnotationMemoryRepo) Update(a *annotation.Annotation) error {
	_, ok := r.annotations[a.ID().ID]
	if !ok {
		return annotation.ErrNotUpdated
	}
	r.annotations[a.ID().ID] = a
	return nil
}

func (r *AnnotationMemoryRepo) Delete(id string) error {
	aid, err := uuid.Parse(id)
	if err != nil {
		return annotation.ErrNotDeleted
	}

	_, ok := r.annotations[aid]
	if !ok {
		return annotation.ErrNotDeleted
	}
	delete(r.annotations, aid)
	return nil
}
