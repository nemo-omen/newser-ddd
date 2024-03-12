package value

import domainerror "newser/internal/domain/error"

type Term = string

func NewTerm(str string) (Term, error) {
	if len(str) < 1 {
		return "", domainerror.ErrInvalidValue
	}
	return str, nil
}
