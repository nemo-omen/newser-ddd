package entity

import (
	"newser/internal/domain/value"
)

type CategoryId = value.ID

type Category struct {
	ID   CategoryId
	Term string
}
