package query

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain"
)

type ExistSeller struct {
	repository domain.ISellerRepository
}

func NewExistSellerCommand(repository domain.ISellerRepository) *ExistSeller {
	return &ExistSeller{repository: repository}
}

func (queryCommand ExistSeller) Execute(id uuid.UUID) error {
	_, err := queryCommand.repository.FindById(id)
	if err != nil {
		return err
	}

	return nil
}
