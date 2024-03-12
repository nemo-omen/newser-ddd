package user

import (
	"encoding/json"
	"newser/internal/domain/entity"
	"newser/internal/domain/value"
)

type User struct {
	user          *entity.Person
	subscriptions []entity.SubscriptionId
	annotations   []entity.NoteId
	collections   []entity.CollectionId
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
		ID:    value.NewId(),
		Name:  userName,
		Email: userEmail,
	}

	user := User{
		user:          person,
		subscriptions: []entity.SubscriptionId{},
		annotations:   []entity.NoteId{},
		collections:   []entity.CollectionId{},
	}
	return &user, nil
}

func (u *User) JSON() []byte {
	type jsonUser struct {
		ID            value.ID                `json:"id"`
		Name          value.Name              `json:"name"`
		Email         value.Email             `json:"email"`
		Subscriptions []entity.SubscriptionId `json:"subscriptions"`
		Annotations   []entity.NoteId         `json:"annotations"`
		Collections   []entity.CollectionId   `json:"collections"`
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
