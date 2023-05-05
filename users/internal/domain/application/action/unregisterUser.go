package action

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/model"
)

type UnregisterUserEvent struct {
	repository model.IUserRepository
}

func NewUnregisterUserEvent(repository model.IUserRepository) *UnregisterUserEvent {
	return &UnregisterUserEvent{
		repository: repository,
	}
}

func (actionEvent UnregisterUserEvent) Execute(ID uuid.UUID) error {
	return actionEvent.repository.Delete(ID)
}
