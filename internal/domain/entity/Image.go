package entity

import (
	"newser/internal/domain/value"
)

type ImageId = value.ID

type Image struct {
	ID    ImageId
	Title string
	URL   string
}
