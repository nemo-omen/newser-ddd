package subscription

import (
	"newser/internal/domain/entity"
)

type SubscriptionAggregate struct {
	subscription entity.Subscription
	user         entity.PersonId
	feed         entity.FeedId
	// created ?
}
