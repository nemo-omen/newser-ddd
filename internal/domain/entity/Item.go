package entity

import (
	"time"

	"newser/internal/domain/value"
)

type ItemId = value.ID

type Item struct {
	ID              ItemId
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
