package entity

import (
	"newser/internal/domain/value"
)

type Person struct {
	ID    value.PersonId
	Name  string
	Email string
}
