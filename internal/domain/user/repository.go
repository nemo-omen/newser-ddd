package user

type UserRepository interface {
	Save(user *User) error
	FindById(id string) (*User, error)
	FindByEmail(email string) (*User, error)
	FindAll() ([]*User, error)
	Update(user *User) error
	Delete(id string) error
}
