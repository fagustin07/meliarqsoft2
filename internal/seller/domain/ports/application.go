package ports

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain"
)

type ISellerService interface {
	Create(businessName string, email string) (uuid.UUID, error)
	Update(id uuid.UUID, businessName string, email string) error
	Delete(ID uuid.UUID) error
	Find(name string) ([]*domain.Seller, error)
	Exist(seller uuid.UUID) error
}
