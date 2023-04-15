package ports

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/seller/domain"
)

type ISellerRepository interface {
	Create(seller *domain.Seller) error
	Update(id uuid.UUID, businessName string, email string) error
	Delete(ID uuid.UUID) error
	Find(businessName string) ([]*domain.Seller, error)
	FindById(idSeller uuid.UUID) (*domain.Seller, error)
}
