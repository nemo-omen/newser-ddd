package entity

import (
	"encoding/json"
	"newser/internal/domain/value"
)

type Person struct {
	ID    value.PersonId
	Name  string
	Email string
}

type PersonProps struct {
	Name  string
	Email string
}

func NewPerson(props PersonProps) (*Person, error) {
	validName, err := value.NewName(props.Name)
	if err != nil {
		return nil, err
	}

	validEmail, err := value.NewEmail(props.Email)
	if err != nil {
		return nil, err
	}

	return &Person{
		ID:    value.PersonId{ID: value.NewId()},
		Name:  validName,
		Email: validEmail,
	}, nil
}

func (p *Person) JSON() []byte {
	type jsonPerson struct {
		ID    value.PersonId `json:"id"`
		Name  string         `json:"name"`
		Email string         `json:"email"`
	}
	jp := jsonPerson{
		ID:    p.ID,
		Name:  p.Name,
		Email: p.Email,
	}
	b, _ := json.MarshalIndent(jp, "", "  ")
	return b
}

func (p *Person) String() string {
	return string(p.JSON())
}
