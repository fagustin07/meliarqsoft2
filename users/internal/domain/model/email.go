package model

import (
	"meliarqsoft2/pkg/exceptions/model"
	"net/mail"
)

type Email struct {
	Address string
}

func NewEmail(address string) (*Email, error) {
	_, err := mail.ParseAddress(address)

	if err != nil {
		return nil, model.InvalidEmailError{}
	} else {
		return &Email{Address: address}, nil
	}
}
