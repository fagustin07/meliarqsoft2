package query

import (
	"meliarqsoft2/internal/domain"
)

type FindUserEvent struct {
	repository domain.IUserRepository
}

func NewFindUserEvent(repository domain.IUserRepository) *FindUserEvent {
	return &FindUserEvent{
		repository: repository,
	}
}
func (queryEvent FindUserEvent) Execute(emailPattern string) ([]*domain.User, error) {
	return queryEvent.repository.Find(emailPattern)
}
