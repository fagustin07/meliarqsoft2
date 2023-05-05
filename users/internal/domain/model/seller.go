package model

import (
	"github.com/google/uuid"
)

type Seller struct {
	ID           uuid.UUID
	BusinessName string
	Email        *Email
}

//go:generate mockgen -destination=../mock/sellerRepository.go -package=mock -source=seller.go
type ISellerRepository interface {
	Create(seller *Seller) (uuid.UUID, error)
	Update(id uuid.UUID, businessName string, email string) error
	Delete(ID uuid.UUID) error
	Find(businessName string) ([]*Seller, error)
	FindById(idSeller uuid.UUID) (*Seller, error)
	FindByEmail(email string) (*Seller, error)
	FindByBusinessName(businessName string) (*Seller, error)
}