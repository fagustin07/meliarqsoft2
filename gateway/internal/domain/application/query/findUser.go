package query

import (
	"meliarqsoft2/internal/domain/model"
)

type FindUserEvent struct {
	repository model.IUserRepository
}

func NewFindUserEvent(repository model.IUserRepository) *FindUserEvent {
	return &FindUserEvent{
		repository: repository,
	}
}
func (queryEvent FindUserEvent) Execute(emailPattern string) ([]*model.User, error) {
	return queryEvent.repository.Find(emailPattern)
}
