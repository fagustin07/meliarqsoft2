package query

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/user/domain/ports"
)

type ExistUser struct {
	repository ports.IUserRepository
}

func NewExistUserCommand(repository ports.IUserRepository) *ExistUser {
	return &ExistUser{repository: repository}
}

func (queryCommand ExistUser) Execute(userID uuid.UUID) error {
	_, err := queryCommand.repository.FindById(userID)
	if err != nil {
		return err
	}

	return nil
}
