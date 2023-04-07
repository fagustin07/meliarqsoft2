package ports

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/seller/domain"
)

type ISellerManager interface {
	Create(businessName string, email string) (uuid.UUID, error)
	Update(id uuid.UUID, businessName string, email string) error
	Delete(ID uuid.UUID) error
	Find(name string) ([]*domain.Seller, error)
}
