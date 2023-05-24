package action

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/model"
)

type RegisterCustomerEvent struct {
	repository          model.ICustomerRepository
	notificationService model.INotificationService
}

func NewRegisterCustomerEvent(repository model.ICustomerRepository, service model.INotificationService) *RegisterCustomerEvent {
	return &RegisterCustomerEvent{
		repository:          repository,
		notificationService: service,
	}
}

func (actionEvent RegisterCustomerEvent) Execute(customer *model.Customer, notification *model.Notification) (uuid.UUID, error) {
	id, err := actionEvent.repository.Create(customer)
	if err != nil {
		return uuid.Nil, err
	}

	_ = actionEvent.notificationService.Send(notification)

	return id, nil
}
