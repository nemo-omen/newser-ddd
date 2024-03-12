package entity

import (
	"newser/internal/domain/value"
)

type Person struct {
	ID    value.ID
	Name  string
	Email string
}
