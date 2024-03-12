package entity

import (
	"newser/internal/domain/value"
)

type Image struct {
	ID    value.ImageId
	Title string
	URL   string
}
