package domain

import "github.com/google/uuid"

type ISellerRepository interface {
	Create(ID uuid.UUID, businessName string, email string) error
}
