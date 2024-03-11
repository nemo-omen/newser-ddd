package article

import (
	"newser/internal/domain/entity"
)

type Article struct {
	item        *entity.Item
	authors     []entity.PersonId
	image       entity.ImageId
	annotations []entity.NoteId
	feed        entity.FeedId
	slug        string
}
