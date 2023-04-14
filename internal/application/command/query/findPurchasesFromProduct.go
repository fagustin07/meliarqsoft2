package query

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/purchase/domain"
	"meliarqsoft2/internal/purchase/domain/ports"
)

type FindPurchasesFromProduct struct {
	repository ports.IPurchaseRepository
}

func (queryCommand FindPurchasesFromProduct) Execute(productID uuid.UUID) ([]*domain.Purchase, error) {
	return queryCommand.repository.Find(productID)
}
