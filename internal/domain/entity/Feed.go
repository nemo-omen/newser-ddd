package entity

import (
	"newser/internal/domain/value"
)

type FeedId = value.ID

type Feed struct {
	ID          FeedId
	Title       value.Title
	SiteUrl     value.Link
	FeedUrl     value.Link
	Description value.Description
	Copyright   string
	Language    string
	FeedType    string
}

type FeedProps struct {
	Title       string
	SiteUrl     string
	FeedUrl     string
	Description string
	Copyright   string
	Language    string
	FeedType    string
}

func NewFeed(props FeedProps) (*Feed, error) {
	title, err := value.NewTitle(props.Title)
	if err != nil {
		return nil, err
	}

	siteUrl, err := value.NewLink(props.SiteUrl)
	if err != nil {
		return nil, err
	}

	feedUrl, err := value.NewLink(props.FeedUrl)
	if err != nil {
		return nil, err
	}

	desc := value.NewDescription(props.Description)
}
