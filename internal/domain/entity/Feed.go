package entity

import (
	"newser/internal/domain/value"
)

type FeedId = value.ID

type Feed struct {
	ID          FeedId
	Title       string
	SiteUrl     string
	FeedUrl     string
	Description string
	Copyright   string
	Language    string
	FeedType    string
}
