package domain

import "github.com/google/uuid"

type Seller struct {
	ID           uuid.UUID
	BusinessName string
	Email        string
}

func NewSeller(id uuid.UUID, businessName string, email string) *Seller {
	return &Seller{ID: id, BusinessName: businessName, Email: email}
}
