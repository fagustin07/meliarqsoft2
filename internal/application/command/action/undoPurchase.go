package action

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/purchase/domain/ports"
)

type UndoPurchase struct {
	repository ports.IPurchaseRepository
}

func NewUndoPurchaseCommand(repo ports.IPurchaseRepository) *UndoPurchase {
	return &UndoPurchase{
		repository: repo,
	}
}

func (command UndoPurchase) Execute(id uuid.UUID) error {
	err := command.repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
