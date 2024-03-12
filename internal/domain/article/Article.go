package article

import (
	"newser/internal/domain/entity"
	"newser/internal/domain/value"
)

type Article struct {
	item        *entity.Item
	authors     []entity.PersonId
	image       entity.ImageId
	annotations []entity.NoteId
	categories  []entity.CategoryId
	feed        entity.FeedId
	slug        string
	read        bool
	saved       bool
}

type ArticleProps struct {
	entity.ItemProps
	Image  string
	FeedId string
}

func NewArticle(props ArticleProps) (*Article, error) {
	item, err := entity.NewItem(props.ItemProps)
	if err != nil {
		return nil, err
	}

	slug, err := value.NewSlug(item.Title)
	if err != nil {
		return nil, err
	}

	article := &Article{
		item:        item,
		authors:     []entity.PersonId{},
		image:       value.IdFromString(props.Image),
		annotations: []entity.NoteId{},
		categories:  []entity.CategoryId{},
		feed:        value.IdFromString(props.FeedId),
		slug:        slug,
	}

	return article, nil
}
