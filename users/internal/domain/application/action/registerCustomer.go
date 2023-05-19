package action

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/model"
)

type RegisterCustomerEvent struct {
	repository model.ICustomerRepository
}

func NewRegisterCustomerEvent(repository model.ICustomerRepository) *RegisterCustomerEvent {
	return &RegisterCustomerEvent{repository: repository}
}

func (actionEvent RegisterCustomerEvent) Execute(customer *model.Customer) (uuid.UUID, error) {
	id, err := actionEvent.repository.Create(customer)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}
