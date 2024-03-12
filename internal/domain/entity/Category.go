package entity

import (
	"newser/internal/domain/value"
)

type Category struct {
	ID   value.ID
	Term string
}
