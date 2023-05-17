package model

import (
	"github.com/google/uuid"
)

type User struct {
	ID      uuid.UUID
	Name    string
	Surname string
	Email   *Email
}

type Customer struct {
	ID      string `json:"user_id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
}

//go:generate mockgen -destination=../mock/userRepository.go -package=mock -source=user.go
type IUserRepository interface {
	Create(user *User) (uuid.UUID, error)
	Update(ID uuid.UUID, name string, surname string, email string) error
	Delete(ID uuid.UUID) error
	Find(emailPattern string) ([]*User, error)
	FindById(idUser uuid.UUID) (*User, error)
}

type IUserService interface {
	Register(user *User) (uuid.UUID, error)
	Update(ID uuid.UUID, name string, surname string, email string) error
	Delete(ID uuid.UUID) error
	Find(emailPattern string) ([]Customer, error)
}
