package action

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain"
	"meliarqsoft2/internal/user/domain/ports"
)

type UpdateUserEvent struct {
	repository ports.IUserRepository
}

func NewUpdateUserEvent(repository ports.IUserRepository) *UpdateUserEvent {
	return &UpdateUserEvent{
		repository: repository,
	}
}

func (actionEvent UpdateUserEvent) Execute(ID uuid.UUID, name string, surname string, email string) error {
	newEmail, err := domain.NewEmail(email)
	if err != nil {
		return err
	}

	return actionEvent.repository.Update(ID, name, surname, newEmail.Address)
}
