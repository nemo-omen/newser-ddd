package subscription

import (
	"encoding/json"
	"newser/internal/domain/value"
)

type Subscription struct {
	id   value.SubscriptionId
	user value.PersonId
	feed value.NewsfeedId
	// created ?
}

type SubscriptionProps struct {
	User value.PersonId
	Feed value.NewsfeedId
}

func NewSubscription(props SubscriptionProps) (*Subscription, error) {
	return &Subscription{
		id:   value.SubscriptionId{ID: value.NewId()},
		user: props.User,
		feed: props.Feed,
	}, nil
}

func (s *Subscription) JSON() []byte {
	type jsonSubscription struct {
		ID   value.SubscriptionId `json:"id"`
		User value.PersonId       `json:"user"`
		Feed value.NewsfeedId     `json:"feed"`
	}
	js := jsonSubscription{
		ID:   s.id,
		User: s.user,
		Feed: s.feed,
	}
	b, _ := json.MarshalIndent(js, "", "  ")
	return b
}

func (s *Subscription) String() string {
	return string(s.JSON())
}
