package value

import (
	"github.com/google/uuid"
)

type ID = uuid.UUID

func NewId() ID {
	return uuid.New()
}
