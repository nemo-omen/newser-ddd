package value

import (
	domainerror "newser/internal/domain/error"

	"github.com/google/uuid"
)

type ID = uuid.UUID

type AnnotationId struct {
	ID
}

type ArticleId struct {
	ID
}

type CategoryId struct {
	ID
}

type CollectionId struct {
	ID
}

type ImageId struct {
	ID
}

type NewsfeedId struct {
	ID
}

type PersonId struct {
	ID
}

type SubscriptionId struct {
	ID
}

func NewId() ID {
	return uuid.New()
}

func IdFromString(id string) (ID, error) {
	u, err := uuid.Parse(id)
	if err != nil {
		return uuid.Nil, domainerror.ErrInvalidValue
	}
	return u, nil
}
