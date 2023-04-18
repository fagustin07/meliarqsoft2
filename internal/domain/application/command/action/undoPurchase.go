package action

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain"
)

type UndoPurchase struct {
	repository domain.IPurchaseRepository
}

func NewUndoPurchaseCommand(repo domain.IPurchaseRepository) *UndoPurchase {
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
