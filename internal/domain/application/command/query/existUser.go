package query

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain"
)

type ExistUser struct {
	repository domain.IUserRepository
}

func NewExistUserCommand(repository domain.IUserRepository) *ExistUser {
	return &ExistUser{repository: repository}
}

func (queryCommand ExistUser) Execute(userID uuid.UUID) error {
	_, err := queryCommand.repository.FindById(userID)
	if err != nil {
		return err
	}

	return nil
}
