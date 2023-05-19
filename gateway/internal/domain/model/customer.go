package model

import (
	"github.com/google/uuid"
)

type Customer struct {
	ID      uuid.UUID `json:"customer_id"`
	Name    string    `json:"name"`
	Surname string    `json:"surname"`
	Email   string    `json:"email"`
}

type CreateCustomerRequest struct {
	Name    string `json:"name" bson:"name" `
	Surname string `json:"surname" bson:"surname" `
	Email   string `json:"email" bson:"email" `
}

type UpdateCustomerRequest struct {
	Name    string `json:"name" bson:"name" `
	Surname string `json:"surname" bson:"surname" `
	Email   string `json:"email" bson:"email"`
}

type CustomerID struct {
	ID uuid.UUID `json:"customer_id" bson:"customer_id"`
}

type ICustomerService interface {
	Register(customer CreateCustomerRequest) (uuid.UUID, error)
	Update(ID uuid.UUID, name string, surname string, email string) error
	Delete(ID uuid.UUID) error
	Find(emailPattern string) ([]Customer, error)
}
