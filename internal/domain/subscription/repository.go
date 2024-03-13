package subscription

import "errors"

var (
	ErrNotFound   = errors.New("subscription not found")
	ErrNotSaved   = errors.New("subscription not saved")
	ErrNotUpdated = errors.New("subscription not updated")
	ErrNotDeleted = errors.New("subscription not deleted")
)

type SubscriptionRepository interface {
	Save(subscription *Subscription) error
	FindById(id string) (*Subscription, error)
	FindByUserId(userId string) ([]*Subscription, error)
	FindByNewsfeedId(feedId string) ([]*Subscription, error)
	FindByNewsfeedIdAndUserId(feedId string, userId string) (*Subscription, error)
	Delete(id string) error
}
