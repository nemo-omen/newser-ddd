package entity

import (
	"newser/internal/domain/value"
)

type PersonId = value.ID

type Person struct {
	ID    PersonId
	Name  string
	Email string
}
