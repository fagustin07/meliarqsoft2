package query

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/model"
)

type ExistCustomer struct {
	repository model.ICustomerRepository
}

func NewExistCustomerCommand(repository model.ICustomerRepository) *ExistCustomer {
	return &ExistCustomer{repository: repository}
}

func (queryCommand ExistCustomer) Execute(customerID uuid.UUID) error {
	_, err := queryCommand.repository.FindById(customerID)
	if err != nil {
		return err
	}

	return nil
}
