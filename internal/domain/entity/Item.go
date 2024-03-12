package entity

import (
	"time"

	"newser/internal/domain/value"
)

type ItemId = value.ID
type ItemTitle = value.Title

type Item struct {
	ID              ItemId
	Title           ItemTitle
	Description     value.Description
	Content         value.ItemContent
	Link            value.Link
	Links           []value.Link
	Updated         string
	UpdatedParsed   time.Time
	Published       string
	PublishedParsed time.Time
	GUID            string
}

type ItemProps struct {
	Title           string
	Description     string
	Content         string
	Link            string
	Links           []string
	Updated         string
	UpdatedParsed   time.Time
	Published       string
	PublishedParsed time.Time
	GUID            string
}

func NewItem(props ItemProps) (*Item, error) {
	validTitle, err := value.NewTitle(props.Title)
	if err != nil {
		return nil, err
	}
	// sanitize content & add 'target="_blank"' to links
	preparedContent := value.NewItemContent(props.Content)

	desc := ""
	// set description to passed item description first
	if len(props.Description) > 0 {
		desc = props.Description
	}

	// if item content is long enough
	// replace description
	if len(preparedContent) > 240 {
		desc = preparedContent[:240] + "..."
	}

	strippedDesc := value.NewDescription(desc)

	l, _ := value.NewLink(props.Link)
	ll := []value.Link{}

	for _, href := range props.Links {
		lnk, _ := value.NewLink(href)
		if len(lnk) > 0 {
			ll = append(ll, lnk)
		}
	}

	item := &Item{
		Title:           validTitle,
		Description:     strippedDesc,
		Content:         preparedContent,
		Link:            l,
		Links:           ll,
		Updated:         props.Updated,
		UpdatedParsed:   props.UpdatedParsed,
		Published:       props.Published,
		PublishedParsed: props.PublishedParsed,
		GUID:            props.GUID,
	}

	return item, nil
}
