package domain

import "github.com/google/uuid"

type User struct {
	ID      uuid.UUID
	Name    string
	Surname string
	Email   string
}

func NewUser(id uuid.UUID, name string, surname string, email string) *User {
	return &User{ID: id, Name: name, Surname: surname, Email: email}
}
