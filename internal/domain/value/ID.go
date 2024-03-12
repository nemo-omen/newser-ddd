package value

import (
	"github.com/google/uuid"
)

type ID = uuid.UUID

func NewId() ID {
	return uuid.New()
}

func IdFromString(id string) ID {
	u, _ := uuid.Parse(id)
	return u
}
