package query

import (
	"meliarqsoft2/internal/user/domain"
	"meliarqsoft2/internal/user/domain/ports"
)

type FindUser struct {
	repository ports.IUserRepository
}

func (queryEvent FindUser) Execute(emailPattern string) ([]*domain.User, error) {
	return queryEvent.repository.Find(emailPattern)
}
