package model

import (
	"meliarqsoft2/pkg/exceptions/model"
	"net/mail"
)

type Notification struct {
	Name  string
	Email string
	Key   string
}

func NewNotification(name string, email string) (*Notification, error) {
	_, err := mail.ParseAddress(email)

	if err != nil {
		return nil, model.InvalidEmailError{}
	} else {
		return &Notification{
			Name:  name,
			Email: email,
			Key:   "Register",
		}, nil
	}
}

type INotificationRepository interface {
	Send(notification *Notification) error
}
