package model

import (
	"github.com/google/uuid"
)

type Customer struct {
	ID      uuid.UUID
	Name    string
	Surname string
	Email   *Email
}

//go:generate mockgen -destination=../mock/customerRepository.go -package=mock -source=customer.go
type ICustomerRepository interface {
	Create(customer *Customer) (uuid.UUID, error)
	Update(ID uuid.UUID, name string, surname string, email string) error
	Delete(ID uuid.UUID) error
	Find(emailPattern string) ([]*Customer, error)
	FindById(idCustomer uuid.UUID) (*Customer, error)
}
