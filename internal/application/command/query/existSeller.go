package query

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/seller/domain/ports"
)

type ExistSeller struct {
	repository ports.ISellerRepository
}

func NewExistSellerCommand(repository ports.ISellerRepository) *ExistSeller {
	return &ExistSeller{repository: repository}
}

func (queryCommand ExistSeller) Execute(id uuid.UUID) error {
	_, err := queryCommand.repository.FindById(id)
	if err != nil {
		return err
	}

	return nil
}
