package action

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain"
)

type UnregisterUserEvent struct {
	repository domain.IUserRepository
}

func NewUnregisterUserEvent(repository domain.IUserRepository) *UnregisterUserEvent {
	return &UnregisterUserEvent{
		repository: repository,
	}
}

func (actionEvent UnregisterUserEvent) Execute(ID uuid.UUID) error {
	return actionEvent.repository.Delete(ID)
}
