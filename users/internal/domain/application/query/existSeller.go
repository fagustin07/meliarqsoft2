package query

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/model"
)

type ExistSeller struct {
	repository model.ISellerRepository
}

func NewExistSellerCommand(repository model.ISellerRepository) *ExistSeller {
	return &ExistSeller{repository: repository}
}

func (queryCommand ExistSeller) Execute(id uuid.UUID) error {
	_, err := queryCommand.repository.FindById(id)
	if err != nil {
		return err
	}

	return nil
}
