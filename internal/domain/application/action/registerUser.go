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

func (actionEvent RegisterUserEvent) Execute(user *model.User) (uuid.UUID, error) {
	id, err := actionEvent.repository.Create(user)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}
