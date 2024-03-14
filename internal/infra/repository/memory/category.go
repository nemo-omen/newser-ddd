package memory

import (
	"newser/internal/domain/category"
	"sync"

	"github.com/google/uuid"
)

type CategoryMemoryRepo struct {
	mu         sync.RWMutex
	categories map[uuid.UUID]*category.Category
}

func NewCategoryMemoryRepo() *CategoryMemoryRepo {
	return &CategoryMemoryRepo{
		categories: make(map[uuid.UUID]*category.Category),
	}
}

func (r *CategoryMemoryRepo) Save(c *category.Category) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.categories[c.ID.ID] = c
	return nil
}

func (r *CategoryMemoryRepo) FindById(id uuid.UUID) (*category.Category, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	c, ok := r.categories[id]
	if !ok {
		return nil, category.ErrNotFound
	}
	return c, nil
}

func (r *CategoryMemoryRepo) FindByTerm(term string) (*category.Category, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, c := range r.categories {
		if c.Term == term {
			return c, nil
		}
	}
	return nil, category.ErrNotFound
}

func (r *CategoryMemoryRepo) FindAll() ([]*category.Category, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	categories := make([]*category.Category, 0, len(r.categories))
	for _, c := range r.categories {
		categories = append(categories, c)
	}
	return categories, nil
}
