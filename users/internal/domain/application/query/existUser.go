package query

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/model"
)

type ExistUser struct {
	repository model.IUserRepository
}

func NewExistUserCommand(repository model.IUserRepository) *ExistUser {
	return &ExistUser{repository: repository}
}

func (queryCommand ExistUser) Execute(userID uuid.UUID) error {
	_, err := queryCommand.repository.FindById(userID)
	if err != nil {
		return err
	}

	return nil
}
