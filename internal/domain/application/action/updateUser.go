package action

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/model"
)

type UpdateUserEvent struct {
	repository model.IUserRepository
}

func NewUpdateUserEvent(repository model.IUserRepository) *UpdateUserEvent {
	return &UpdateUserEvent{
		repository: repository,
	}
}

func (actionEvent UpdateUserEvent) Execute(ID uuid.UUID, name string, surname string, email string) error {
	newEmail, err := model.NewEmail(email)
	if err != nil {
		return err
	}

	return actionEvent.repository.Update(ID, name, surname, newEmail.Address)
}
