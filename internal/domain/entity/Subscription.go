package entity

import (
	"newser/internal/domain/value"
)

type SubscriptionId = value.ID

type Subscription struct {
	ID value.ID
}
