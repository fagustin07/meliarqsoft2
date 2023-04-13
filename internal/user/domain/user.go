package domain

import (
	"errors"
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain"
)

type User struct {
	ID      uuid.UUID
	Name    string
	Surname string
	Email   *domain.Email
}

func NewUser(id uuid.UUID, name string, surname string, email string) (*User, error) {
	newEmail, err := domain.NewEmail(email)
	if err != nil {
		return nil, errors.New("invalid email")
	}

	return &User{ID: id, Name: name, Surname: surname, Email: newEmail}, nil
}
