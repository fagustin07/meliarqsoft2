package ports

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/purchase/domain"
)

type IPurchaseManager interface {
	Create(IDProduct uuid.UUID, IDUser uuid.UUID, units int) (*domain.Purchase, error)
	Find(productID uuid.UUID) ([]*domain.Purchase, error)
}
