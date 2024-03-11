package collection

import (
	"newser/internal/domain/entity"
)

type Collection struct {
	collection entity.Set
	user       entity.PersonId
	items      []entity.ItemId
}
