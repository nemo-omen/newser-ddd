package subscription

import (
	"newser/internal/domain/entity"
)

type UserSubscription struct {
	subscription *entity.Subscription
	user         entity.PersonId
	feed         entity.FeedId
	// created ?
}
