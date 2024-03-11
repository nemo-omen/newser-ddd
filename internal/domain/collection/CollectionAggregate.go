package collection

import (
	"newser/internal/domain/entity"
)

type CollectionAggregate struct {
	collection entity.Collection
	user       entity.PersonId
	items      []entity.ItemId
}
