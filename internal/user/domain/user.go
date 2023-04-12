package domain

import (
	"errors"
	"github.com/google/uuid"
	"meliarqsoft2/internal/helpers/value_objects"
)

type User struct {
	ID      uuid.UUID
	Name    string
	Surname string
	Email   *value_objects.Email
}

func NewUser(id uuid.UUID, name string, surname string, email string) (*User, error) {
	newEmail, err := value_objects.NewEmail(email)
	if err != nil {
		return nil, errors.New("invalid email")
	}

	return &User{ID: id, Name: name, Surname: surname, Email: newEmail}, nil
}
