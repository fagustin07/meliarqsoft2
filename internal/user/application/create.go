package application

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/user/domain"
)

func (service UserService) Create(name string, surname string, email string) (uuid.UUID, error) {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return uuid.Nil, err
	}

	newUser, err := domain.NewUser(newUUID, name, surname, email)
	if err != nil {
		return uuid.Nil, err
	}

	err = service.repository.Create(newUser)
	if err != nil {
		return uuid.Nil, err
	}

	return newUUID, nil
}
