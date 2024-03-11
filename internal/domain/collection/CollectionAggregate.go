package collection

import (
	"newser/internal/domain/entity"
)

type UserCollection struct {
	collection *entity.Collection
	user       entity.PersonId
	items      []entity.ItemId
}
