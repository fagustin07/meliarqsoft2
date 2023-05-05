package action

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/model"
)

type MakePurchaseEvent struct {
	repository model.IPurchaseRepository
}

func NewMakePurchaseEvent(repo model.IPurchaseRepository) *MakePurchaseEvent {
	return &MakePurchaseEvent{
		repository: repo,
	}
}

func (actionEvent MakePurchaseEvent) Execute(purchase *model.Purchase) (uuid.UUID, error) {
	id, err := actionEvent.repository.Create(purchase)

	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}
