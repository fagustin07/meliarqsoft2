package action

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/model"
)

type RegisterUserEvent struct {
	repository model.IUserRepository
}

func NewRegisterUserEvent(repository model.IUserRepository) *RegisterUserEvent {
	return &RegisterUserEvent{repository: repository}
}

func (actionEvent RegisterUserEvent) Execute(name string, surname string, email string) (uuid.UUID, error) {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return uuid.Nil, err
	}

	newUser, err := model.NewUser(newUUID, name, surname, email)
	if err != nil {
		return uuid.Nil, err
	}

	err = actionEvent.repository.Create(newUser)
	if err != nil {
		return uuid.Nil, err
	}

	return newUUID, nil
}
