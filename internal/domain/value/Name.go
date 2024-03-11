package value

import (
	domainerror "newser/internal/domain/error"
)

type Name = string

func NewName(n string) (Name, error) {
	if len(n) < 1 {
		return "", domainerror.ErrInvalidValue
	}
	return n, nil
}
