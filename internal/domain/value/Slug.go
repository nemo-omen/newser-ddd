package value

import (
	domainerror "newser/internal/domain/error"

	"github.com/gosimple/slug"
)

type Slug = string

func NewSlug(str string) (Slug, error) {
	if len(str) < 1 {
		return "", domainerror.ErrInvalidValue
	}
	return slug.Make(str), nil
}
