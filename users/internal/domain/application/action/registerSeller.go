package action

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/model"
)

type RegisterSellerEvent struct {
	repository          model.ISellerRepository
	notificationService model.INotificationService
}

func NewRegisterSellerEvent(repository model.ISellerRepository, service model.INotificationService) *RegisterSellerEvent {
	return &RegisterSellerEvent{
		repository:          repository,
		notificationService: service,
	}
}

func (event RegisterSellerEvent) Execute(seller *model.Seller, notification *model.Notification) (uuid.UUID, error) {
	newId, err := event.repository.Create(seller)
	if err != nil {
		return uuid.Nil, err
	}

	_ = event.notificationService.Send(notification)

	return newId, nil
}
