package query

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/model"
)

type FindPurchasesFromProductEvent struct {
	repository model.IPurchaseRepository
}

func NewFindPurchasesFromProductEvent(repository model.IPurchaseRepository) *FindPurchasesFromProductEvent {
	return &FindPurchasesFromProductEvent{repository: repository}
}

func (queryCommand FindPurchasesFromProductEvent) Execute(productID uuid.UUID) ([]*model.Purchase, error) {
	return queryCommand.repository.Find(productID)
}
