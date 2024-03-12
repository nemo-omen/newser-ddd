package entity

import (
	"newser/internal/domain/value"
)

type Image struct {
	ID    value.ID
	Title string
	URL   string
}
