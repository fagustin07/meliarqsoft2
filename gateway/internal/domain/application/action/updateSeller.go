package action

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/model"
)

type UpdateSellerEvent struct {
	repository model.ISellerRepository
}

func NewUpdateSellerEvent(repository model.ISellerRepository) *UpdateSellerEvent {
	return &UpdateSellerEvent{
		repository: repository,
	}
}

func (actionEvent UpdateSellerEvent) Execute(id uuid.UUID, businessName string, email string) error {
	newEmail, err := model.NewEmail(email)
	if err != nil {
		return err
	}

	return actionEvent.repository.Update(id, businessName, newEmail.Address)
}
