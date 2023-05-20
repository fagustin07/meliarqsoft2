package action

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/model"
)

type UpdateCustomerEvent struct {
	repository model.ICustomerRepository
}

func NewUpdateCustomerEvent(repository model.ICustomerRepository) *UpdateCustomerEvent {
	return &UpdateCustomerEvent{
		repository: repository,
	}
}

func (actionEvent UpdateCustomerEvent) Execute(ID uuid.UUID, name string, surname string, email string) error {
	newEmail, err := model.NewEmail(email)
	if err != nil {
		return err
	}

	return actionEvent.repository.Update(ID, name, surname, newEmail.Address)
}
