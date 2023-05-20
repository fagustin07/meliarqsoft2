package model

import (
	"meliarqsoft2/pkg/exceptions/model"
	"net/mail"
)

type Notification struct {
	Name    string
	Email   string
	Product string
	Key     string
}

func NewNotification(name string, email string, product string, key string) (*Notification, error) {
	_, err := mail.ParseAddress(email)

	if err != nil {
		return nil, model.InvalidEmailError{}
	} else {
		return &Notification{
			Name:    name,
			Email:   email,
			Product: product,
			Key:     key,
		}, nil
	}
}

//go:generate mockgen -destination=../mock/notificationRepository.go -package=mock -source=notification.go
type INotificationRepository interface {
	Send(notification *Notification) error
}
