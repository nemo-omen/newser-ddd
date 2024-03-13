package memory

import (
	"newser/internal/domain/subscription"
	"sync"

	"github.com/google/uuid"
)

type MemorySubscriptionRepository struct {
	mu            sync.RWMutex
	subscriptions map[uuid.UUID]*subscription.Subscription
}

func NewSubscriptionMemoryRepo() *MemorySubscriptionRepository {
	return &MemorySubscriptionRepository{
		subscriptions: make(map[uuid.UUID]*subscription.Subscription),
	}
}

func (r *MemorySubscriptionRepository) Save(s *subscription.Subscription) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.subscriptions[s.ID().ID] = s
	return nil
}

func (r *MemorySubscriptionRepository) FindById(id uuid.UUID) (*subscription.Subscription, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	s, ok := r.subscriptions[id]
	if !ok {
		return nil, subscription.ErrNotFound
	}
	return s, nil
}

func (r *MemorySubscriptionRepository) FindByUserId(userId uuid.UUID) ([]*subscription.Subscription, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	subscriptions := make([]*subscription.Subscription, 0, len(r.subscriptions))
	for _, s := range r.subscriptions {
		if s.User().ID == userId {
			subscriptions = append(subscriptions, s)
		}
	}
	return subscriptions, nil
}

func (r *MemorySubscriptionRepository) FindByNewsfeedId(newsfeedId uuid.UUID) ([]*subscription.Subscription, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	subscriptions := make([]*subscription.Subscription, 0, len(r.subscriptions))
	for _, s := range r.subscriptions {
		if s.Feed().ID == newsfeedId {
			subscriptions = append(subscriptions, s)
		}
	}
	return subscriptions, nil
}

func (r *MemorySubscriptionRepository) FindByNewsfeedIdAndUserId(newsfeedId, userId uuid.UUID) (*subscription.Subscription, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, s := range r.subscriptions {
		if s.Feed().ID == newsfeedId && s.User().ID == userId {
			return s, nil
		}
	}
	return nil, subscription.ErrNotFound
}

func (r *MemorySubscriptionRepository) Delete(id uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.subscriptions, id)
	return nil
}
