package query

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain"
)

type FindPurchasesFromProductEvent struct {
	repository domain.IPurchaseRepository
}

func NewFindPurchasesFromProductEvent(repository domain.IPurchaseRepository) *FindPurchasesFromProductEvent {
	return &FindPurchasesFromProductEvent{repository: repository}
}

func (queryCommand FindPurchasesFromProductEvent) Execute(productID uuid.UUID) ([]*domain.Purchase, error) {
	return queryCommand.repository.Find(productID)
}
