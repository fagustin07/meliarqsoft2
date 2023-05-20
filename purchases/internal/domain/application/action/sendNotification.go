package action

import (
	"meliarqsoft2/internal/domain/model"
)

type SendNotificationEvent struct {
	repository model.INotificationRepository
}

func NewSendNotificationEvent(repository model.INotificationRepository) *SendNotificationEvent {
	return &SendNotificationEvent{
		repository: repository,
	}
}

func (actionEvent SendNotificationEvent) Execute(notification *model.Notification) error {
	return actionEvent.repository.Send(notification)
}
