package ports

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/purchase/domain"
)

type IPurchaseRepository interface {
	Create(purchase *domain.Purchase) error
	Find(productID uuid.UUID) ([]*domain.Purchase, error)
	DeleteMany(productId uuid.UUID) error
}
