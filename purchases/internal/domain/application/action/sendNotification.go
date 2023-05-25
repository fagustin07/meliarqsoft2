package action

import (
	"meliarqsoft2/internal/domain/model"
)

type SendNotificationEvent struct {
	service model.INotificationService
}

func NewSendNotificationEvent(repository model.INotificationService) *SendNotificationEvent {
	return &SendNotificationEvent{
		service: repository,
	}
}

func (actionEvent SendNotificationEvent) Execute(notification *model.Notification) error {
	return actionEvent.service.Send(notification)
}
