package subscription

import (
	"newser/internal/domain/entity"
	"newser/internal/domain/value"
)

type UserSubscription struct {
	subscription *entity.Subscription
	user         value.PersonId
	feed         value.NewsfeedId
	// created ?
}
