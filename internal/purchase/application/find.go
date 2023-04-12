package application

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/purchase/domain"
)

func (service PurchaseService) Find(productID uuid.UUID) ([]*domain.Purchase, error) {
	return service.repository.Find(productID)
}
