package annotation

import (
	"newser/internal/domain/value"
)

type Annotation struct {
	id      value.ID
	title   string
	content string
	article value.ID
	// location
	// created
	// updated
	// collection
}
