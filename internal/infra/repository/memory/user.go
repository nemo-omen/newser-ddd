package memory

import (
	"newser/internal/domain/user"
	"sync"

	"github.com/google/uuid"
)

type MemoryUserRepository struct {
	mu    sync.RWMutex
	users map[uuid.UUID]*user.User
}

func NewUserMemoryRepo() *MemoryUserRepository {
	return &MemoryUserRepository{
		users: make(map[uuid.UUID]*user.User),
	}
}

func (r *MemoryUserRepository) Save(u *user.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.users[u.ID().ID] = u
	return nil
}

func (r *MemoryUserRepository) FindById(id uuid.UUID) (*user.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	u, ok := r.users[id]
	if !ok {
		return nil, user.ErrNotFound
	}
	return u, nil
}

func (r *MemoryUserRepository) FindByEmail(email string) (*user.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, u := range r.users {
		if u.Email() == email {
			return u, nil
		}
	}
	return nil, user.ErrNotFound
}

func (r *MemoryUserRepository) FindAll() ([]*user.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	users := make([]*user.User, 0, len(r.users))
	for _, u := range r.users {
		users = append(users, u)
	}
	return users, nil
}

func (r *MemoryUserRepository) Update(u *user.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.users[u.ID().ID] = u
	return nil
}

func (r *MemoryUserRepository) Delete(id uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.users, id)
	return nil
}
