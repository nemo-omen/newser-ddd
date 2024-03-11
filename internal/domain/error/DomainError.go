package domainerror

import (
	"errors"
)

var (
	ErrInvalidValue   = errors.New("invalid value")
	ErrEntityNotFound = errors.New("entity not found")
)

type Error struct {
	appErr    error
	domainErr error
}

func NewError(domainErr, appErr error) error {
	return Error{
		appErr:    appErr,
		domainErr: domainErr,
	}
}

func (e Error) Error() string {
	return errors.Join(e.domainErr, e.appErr).Error()
}
