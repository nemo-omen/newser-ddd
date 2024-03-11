package newsfeed

import (
	"newser/internal/domain/entity"
)

type Newsfeed struct {
	feed       *entity.Feed
	image      entity.ImageId
	articles   []entity.ItemId
	authors    []entity.PersonId
	categories []entity.CategoryId
	slug       string
}
