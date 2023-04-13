package application

import (
	"errors"
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain"
)

func (service UserService) Update(ID uuid.UUID, name string, surname string, email string) error {
	newEmail, err := domain.NewEmail(email)
	if err != nil {
		return errors.New("invalid email")
	}

	return service.repository.Update(ID, name, surname, newEmail.Address)
}
