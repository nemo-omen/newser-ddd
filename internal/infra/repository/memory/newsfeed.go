package memory

import (
	"newser/internal/domain/newsfeed"
	"sync"

	"github.com/google/uuid"
)

type MemoryNewsfeedRepository struct {
	mu        sync.RWMutex
	newsfeeds map[uuid.UUID]*newsfeed.Newsfeed
}

func NewNewsfeedMemoryRepo() *MemoryNewsfeedRepository {
	return &MemoryNewsfeedRepository{
		newsfeeds: make(map[uuid.UUID]*newsfeed.Newsfeed),
	}
}

func (r *MemoryNewsfeedRepository) Save(n *newsfeed.Newsfeed) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.newsfeeds[n.ID().ID] = n
	return nil
}

func (r *MemoryNewsfeedRepository) FindById(id uuid.UUID) (*newsfeed.Newsfeed, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	n, ok := r.newsfeeds[id]
	if !ok {
		return nil, newsfeed.ErrNotFound
	}
	return n, nil
}

func (r *MemoryNewsfeedRepository) FindAll() ([]*newsfeed.Newsfeed, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	newsfeeds := make([]*newsfeed.Newsfeed, 0, len(r.newsfeeds))
	for _, n := range r.newsfeeds {
		newsfeeds = append(newsfeeds, n)
	}
	return newsfeeds, nil
}

func (r *MemoryNewsfeedRepository) FindByTitle(title string) (*newsfeed.Newsfeed, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, n := range r.newsfeeds {
		if n.Title() == title {
			return n, nil
		}
	}
	return nil, newsfeed.ErrNotFound
}

func (r *MemoryNewsfeedRepository) FindBySlug(slug string) (*newsfeed.Newsfeed, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, n := range r.newsfeeds {
		if n.Slug() == slug {
			return n, nil
		}
	}
	return nil, newsfeed.ErrNotFound
}

func (r *MemoryNewsfeedRepository) Delete(id uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.newsfeeds, id)
	return nil
}
