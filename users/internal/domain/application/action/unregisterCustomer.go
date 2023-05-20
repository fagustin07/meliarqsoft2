package action

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/model"
)

type UnregisterCustomerEvent struct {
	repository model.ICustomerRepository
}

func NewUnregisterCustomerEvent(repository model.ICustomerRepository) *UnregisterCustomerEvent {
	return &UnregisterCustomerEvent{
		repository: repository,
	}
}

func (actionEvent UnregisterCustomerEvent) Execute(ID uuid.UUID) error {
	return actionEvent.repository.Delete(ID)
}
