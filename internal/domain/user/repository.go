package user

import "errors"

var (
	ErrNotFound   = errors.New("user not found")
	ErrNotSaved   = errors.New("user not saved")
	ErrNotUpdated = errors.New("user not updated")
	ErrNotDeleted = errors.New("user not deleted")
)

type UserRepository interface {
	Save(user *User) error
	FindById(id string) (*User, error)
	FindByEmail(email string) (*User, error)
	FindAll() ([]*User, error)
	Update(user *User) error
	Delete(id string) error
}
