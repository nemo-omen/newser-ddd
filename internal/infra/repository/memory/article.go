package memory

import (
	"newser/internal/domain/article"
	"sync"

	"github.com/google/uuid"
)

type MemoryArticleRepository struct {
	mu       sync.RWMutex
	articles map[uuid.UUID]*article.Article
}

func NewArticleMemoryRepo() *MemoryArticleRepository {
	return &MemoryArticleRepository{
		articles: make(map[uuid.UUID]*article.Article),
	}
}

func (r *MemoryArticleRepository) Save(a *article.Article) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.articles[a.ID().ID] = a
	return nil
}

func (r *MemoryArticleRepository) FindById(id uuid.UUID) (*article.Article, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	a, ok := r.articles[id]
	if !ok {
		return nil, article.ErrNotFound
	}
	return a, nil
}

func (r *MemoryArticleRepository) FindByNewsfeedId(newsfeedId uuid.UUID) ([]*article.Article, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	articles := make([]*article.Article, 0, len(r.articles))
	for _, a := range r.articles {
		if a.Feed().ID == newsfeedId {
			articles = append(articles, a)
		}
	}
	return articles, nil
}

func (r *MemoryArticleRepository) FindByTitle(title string) (*article.Article, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, a := range r.articles {
		if a.Title() == title {
			return a, nil
		}
	}
	return nil, article.ErrNotFound
}

func (r *MemoryArticleRepository) FindBySlug(slug string) (*article.Article, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, a := range r.articles {
		if a.Slug() == slug {
			return a, nil
		}
	}
	return nil, article.ErrNotFound
}
