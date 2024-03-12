package value

import (
	"net/url"
	domainerror "newser/internal/domain/error"
)

type Link = string

func NewLink(href string) (Link, error) {
	if len(href) < 1 {
		return "", domainerror.ErrInvalidValue
	}

	u, err := url.ParseRequestURI(href)
	if err != nil {
		return "", domainerror.ErrInvalidValue
	}

	return u.String(), nil
}
