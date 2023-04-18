package action

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain"
)

type RegisterUserEvent struct {
	repository domain.IUserRepository
}

func NewRegisterUserEvent(repository domain.IUserRepository) *RegisterUserEvent {
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
