package domain

import (
	"errors"
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain"
)

type Seller struct {
	ID           uuid.UUID
	BusinessName string
	Email        *domain.Email
}

func NewSeller(id uuid.UUID, businessName string, email string) (*Seller, error) {
	newEmail, err := domain.NewEmail(email)
	if err != nil {
		return nil, errors.New("invalid email")
	}
	return &Seller{ID: id, BusinessName: businessName, Email: newEmail}, nil
}
