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

func (actionEvent SendNotificationEvent) Execute(name string, email string) error {
	newNotification, err := model.NewNotification(name, email)
	if err != nil {
		return err
	}

	return actionEvent.repository.Send(newNotification)
}
