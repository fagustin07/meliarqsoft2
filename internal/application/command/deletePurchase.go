package command

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/purchase/domain/ports"
)

type DeletePurchase struct {
	repository ports.IPurchaseRepository
}

func NewDeletePurchase(repo ports.IPurchaseRepository) *DeletePurchase {
	return &DeletePurchase{
		repository: repo,
	}
}

func (command DeletePurchase) Execute(id uuid.UUID) error {
	err := command.repository.DeleteMany(id)
	if err != nil {
		return err
	}

	return nil
}
