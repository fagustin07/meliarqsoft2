package action

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/user/domain/ports"
)

type UnregisterUserEvent struct {
	repository ports.IUserRepository
}

func NewUnregisterUserEvent(repository ports.IUserRepository) *UnregisterUserEvent {
	return &UnregisterUserEvent{
		repository: repository,
	}
}

func (actionEvent UnregisterUserEvent) Execute(ID uuid.UUID) error {
	return actionEvent.repository.Delete(ID)
}
