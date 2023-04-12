package domain

import (
	"errors"
	"github.com/google/uuid"
	"meliarqsoft2/internal/helpers/value_objects"
)

type Seller struct {
	ID           uuid.UUID
	BusinessName string
	Email        *value_objects.Email
}

func NewSeller(id uuid.UUID, businessName string, email string) (*Seller, error) {
	newEmail, err := value_objects.NewEmail(email)
	if err != nil {
		return nil, errors.New("invalid email")
	}
	return &Seller{ID: id, BusinessName: businessName, Email: newEmail}, nil
}
