package user

import (
	"encoding/json"
	"newser/internal/domain/entity"
	"newser/internal/domain/value"
)

type User struct {
	user          *entity.Person
	subscriptions []value.SubscriptionId
	annotations   []value.AnnotationId
	collections   []value.CollectionId
}

func NewUser(email, name string) (*User, error) {
	userName, err := value.NewName(name)
	if err != nil {
		return nil, err
	}

	userEmail, err := value.NewEmail(email)
	if err != nil {
		return nil, err
	}

	person := &entity.Person{
		ID:    value.PersonId{ID: value.NewId()},
		Name:  userName,
		Email: userEmail,
	}

	user := User{
		user:          person,
		subscriptions: []value.SubscriptionId{},
		annotations:   []value.AnnotationId{},
		collections:   []value.CollectionId{},
	}
	return &user, nil
}

func (u *User) ID() value.PersonId {
	return u.user.ID
}

func (u *User) Name() value.Name {
	return u.user.Name
}

func (u *User) Email() value.Email {
	return u.user.Email
}

func (u *User) JSON() []byte {
	type jsonUser struct {
		ID            value.PersonId         `json:"id"`
		Name          value.Name             `json:"name"`
		Email         value.Email            `json:"email"`
		Subscriptions []value.SubscriptionId `json:"subscriptions"`
		Annotations   []value.AnnotationId   `json:"annotations"`
		Collections   []value.CollectionId   `json:"collections"`
	}
	ju := jsonUser{
		ID:            u.user.ID,
		Name:          u.user.Name,
		Email:         u.user.Email,
		Subscriptions: u.subscriptions,
		Annotations:   u.annotations,
		Collections:   u.collections,
	}
	b, _ := json.MarshalIndent(ju, "", "  ")
	return b
}

func (u *User) String() string {
	j := u.JSON()
	return string(j)
}
