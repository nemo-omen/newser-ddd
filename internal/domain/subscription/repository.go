package subscription

type SubscriptionRepository interface {
	Save(subscription *Subscription) error
	FindById(id string) (*Subscription, error)
	FindByUserId(userId string) ([]*Subscription, error)
	FindByFeedId(feedId string) ([]*Subscription, error)
	FindByFeedIdAndUserId(feedId string, userId string) (*Subscription, error)
	Update(subscription *Subscription) error
	Delete(id string) error
}
