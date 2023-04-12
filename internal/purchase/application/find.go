package application

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/purchase/domain"
)

func (manager PurchaseService) Find(productID uuid.UUID) ([]*domain.Purchase, error) {
	return manager.repository.Find(productID)
}
