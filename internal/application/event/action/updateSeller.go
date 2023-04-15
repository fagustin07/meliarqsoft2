package action

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain"
)

type UpdateSellerEvent struct {
	repository domain.ISellerRepository
}

func NewUpdateSellerEvent(repository domain.ISellerRepository) *UpdateSellerEvent {
	return &UpdateSellerEvent{
		repository: repository,
	}
}

func (actionEvent UpdateSellerEvent) Execute(id uuid.UUID, businessName string, email string) error {
	newEmail, err := domain.NewEmail(email)
	if err != nil {
		return err
	}

	return actionEvent.repository.Update(id, businessName, newEmail.Address)
}
