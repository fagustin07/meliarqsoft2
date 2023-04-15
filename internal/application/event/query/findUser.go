package query

import (
	"meliarqsoft2/internal/user/domain"
	"meliarqsoft2/internal/user/domain/ports"
)

type FindUserEvent struct {
	repository ports.IUserRepository
}

func NewFindUserEvent(repository ports.IUserRepository) *FindUserEvent {
	return &FindUserEvent{
		repository: repository,
	}
}
func (queryEvent FindUserEvent) Execute(emailPattern string) ([]*domain.User, error) {
	return queryEvent.repository.Find(emailPattern)
}
