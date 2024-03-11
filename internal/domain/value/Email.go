package value

import (
	"net/mail"
	domainerror "newser/internal/domain/error"
)

type Email = string

func NewEmail(address string) (string, error) {
	email, err := mail.ParseAddress(address)
	if err != nil {
		return "", domainerror.ErrInvalidValue
	}

	return email.Address, nil
}
