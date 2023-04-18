package model

import (
	"errors"
	"github.com/google/uuid"
)

type Seller struct {
	ID           uuid.UUID
	BusinessName string
	Email        *Email
}

func NewSeller(id uuid.UUID, businessName string, email string) (*Seller, error) {
	newEmail, err := NewEmail(email)
	if err != nil {
		return nil, errors.New("invalid email")
	}
	return &Seller{ID: id, BusinessName: businessName, Email: newEmail}, nil
}

//go:generate mockgen -destination=../mock/sellerRepository.go -package=mock -source=seller.go
type ISellerRepository interface {
	Create(seller *Seller) (uuid.UUID, error)
	Update(id uuid.UUID, businessName string, email string) error
	Delete(ID uuid.UUID) error
	Find(businessName string) ([]*Seller, error)
	FindById(idSeller uuid.UUID) (*Seller, error)
}
