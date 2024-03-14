package service

import (
	"errors"
	"newser/internal/domain/annotation"
	"newser/internal/domain/collection"
	"newser/internal/domain/subscription"
	"newser/internal/domain/user"
	"newser/internal/domain/value"
	"newser/internal/infra/repository/memory"
)

var (
	ErrPasswordTooShort = errors.New("password too short")
)

type UserConfiguration func(us *UserService) error

type UserService struct {
	userRepo         user.UserRepository
	collectionRepo   collection.CollectionRepository
	annotationRepo   annotation.AnnotationRepository
	subscriptionRepo subscription.SubscriptionRepository
}

func NewUserService(configs ...UserConfiguration) (*UserService, error) {
	service := &UserService{}

	for _, config := range configs {
		err := config(service)
		if err != nil {
			return nil, err
		}
	}

	return service, nil
}

func UserServiceWithUserRepository(repo user.UserRepository) UserConfiguration {
	return func(us *UserService) error {
		us.userRepo = repo
		return nil
	}
}

func UserServiceWithMemoryUserRepository() UserConfiguration {
	return func(us *UserService) error {
		repo := memory.NewUserMemoryRepo()
		us.userRepo = repo
		return nil
	}
}

func UserServiceWithMemoryCollectionRepository() UserConfiguration {
	return func(us *UserService) error {
		repo := memory.NewCollectionMemoryRepo()
		us.collectionRepo = repo
		return nil
	}
}

func (us *UserService) SaveUser(u *user.User) error {
	return us.userRepo.Save(u)
}

func (us *UserService) FindAllUsers() ([]*user.User, error) {
	return us.userRepo.FindAll()
}

func (us *UserService) FindUserById(id value.PersonId) (*user.User, error) {
	return us.userRepo.FindById(id.ID)
}

func (us *UserService) FindUserByEmail(email string) (*user.User, error) {
	return us.userRepo.FindByEmail(email)
}

func (us *UserService) Signup(email, name, password string) (*user.User, error) {
	if len(password) < 6 {
		return nil, ErrPasswordTooShort
	}

	u, err := user.NewUser(email, name)
	if err != nil {
		return nil, err
	}

	err = us.userRepo.Save(u)
	if err != nil {
		return nil, err
	}

	return u, nil
}
