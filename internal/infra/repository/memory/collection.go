package memory

import (
	"newser/internal/domain/collection"
	"sync"

	"github.com/google/uuid"
)

type MemoryCollectionRepository struct {
	mu          sync.RWMutex
	collections map[uuid.UUID]*collection.Collection
}

func NewCollectionMemoryRepo() *MemoryCollectionRepository {
	return &MemoryCollectionRepository{
		collections: make(map[uuid.UUID]*collection.Collection),
	}
}

func (r *MemoryCollectionRepository) Save(c *collection.Collection) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.collections[c.ID().ID] = c
	return nil
}

func (r *MemoryCollectionRepository) FindById(id uuid.UUID) (*collection.Collection, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	c, ok := r.collections[id]
	if !ok {
		return nil, collection.ErrNotFound
	}
	return c, nil
}

func (r *MemoryCollectionRepository) FindBySlug(slug string) (*collection.Collection, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, c := range r.collections {
		if c.Slug() == slug {
			return c, nil
		}
	}
	return nil, collection.ErrNotFound
}

func (r *MemoryCollectionRepository) FindAllByUserId(userId uuid.UUID) ([]*collection.Collection, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	collections := make([]*collection.Collection, 0, len(r.collections))
	for _, c := range r.collections {
		if c.User().String() == userId.String() {
			collections = append(collections, c)
		}
	}
	return collections, nil
}

func (r *MemoryCollectionRepository) Update(c *collection.Collection) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.collections[c.ID().ID] = c
	return nil
}

func (r *MemoryCollectionRepository) Delete(id uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.collections, id)
	return nil
}
