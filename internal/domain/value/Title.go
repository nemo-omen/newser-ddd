package value

import domainerror "newser/internal/domain/error"

type Title = string

func NewTitle(title string) (Title, error) {
	if len(title) < 1 {
		return "", domainerror.ErrInvalidValue
	}
	return title, nil
}
