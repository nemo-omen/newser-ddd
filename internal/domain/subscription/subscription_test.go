package subscription

import (
	"testing"

	"newser/internal/domain/value"

	"github.com/stretchr/testify/assert"
)

func TestNewSubscription(t *testing.T) {
	userID := value.PersonId{ID: value.NewId()}
	feedID := value.NewsfeedId{ID: value.NewId()}

	subscription, err := NewSubscription(SubscriptionProps{
		User: userID,
		Feed: feedID,
	})

	assert.NoError(t, err)
	assert.Equal(t, userID, subscription.user)
	assert.Equal(t, feedID, subscription.feed)
}
