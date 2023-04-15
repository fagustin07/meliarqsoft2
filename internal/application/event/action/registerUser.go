package action

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/user/domain"
	"meliarqsoft2/internal/user/domain/ports"
)

type RegisterUserEvent struct {
	repository ports.IUserRepository
}

func NewRegisterUserEvent(repository ports.IUserRepository) *RegisterUserEvent {
	return &RegisterUserEvent{repository: repository}
}

func (actionEvent RegisterUserEvent) Execute(name string, surname string, email string) (uuid.UUID, error) {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return uuid.Nil, err
	}

	newUser, err := domain.NewUser(newUUID, name, surname, email)
	if err != nil {
		return uuid.Nil, err
	}

	err = actionEvent.repository.Create(newUser)
	if err != nil {
		return uuid.Nil, err
	}

	return newUUID, nil
}
