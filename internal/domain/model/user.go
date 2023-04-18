package model

import (
	"errors"
	"github.com/google/uuid"
)

type User struct {
	ID      uuid.UUID
	Name    string
	Surname string
	Email   *Email
}

func NewUser(id uuid.UUID, name string, surname string, email string) (*User, error) {
	newEmail, err := NewEmail(email)
	if err != nil {
		return nil, errors.New("invalid email")
	}

	return &User{ID: id, Name: name, Surname: surname, Email: newEmail}, nil
}

//go:generate mockgen -destination=../mock/userRepository.go -package=mock -source=user.go
type IUserRepository interface {
	Create(user *User) error
	Update(ID uuid.UUID, name string, surname string, email string) error
	Delete(ID uuid.UUID) error
	Find(emailPattern string) ([]*User, error)
	FindById(idUser uuid.UUID) (*User, error)
}
