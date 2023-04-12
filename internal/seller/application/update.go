package application

import (
	"errors"
	"github.com/google/uuid"
	"meliarqsoft2/internal/helpers/value_objects"
)

func (service SellerService) Update(id uuid.UUID, businessName string, email string) error {
	newEmail, err := value_objects.NewEmail(email)
	if err != nil {
		return errors.New("invalid email")
	}

	return service.repo.Update(id, businessName, newEmail.Address)
}
