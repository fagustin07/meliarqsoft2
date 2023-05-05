package action

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/model"
)

type UndoPurchase struct {
	repository model.IPurchaseRepository
}

func NewUndoPurchaseCommand(repo model.IPurchaseRepository) *UndoPurchase {
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
