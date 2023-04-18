package model

import (
	"errors"
	"net/mail"
)

type Email struct {
	Address string
}

func NewEmail(address string) (*Email, error) {
	_, err := mail.ParseAddress(address)

	if err != nil {
		return nil, errors.New("invalid email address")
	} else {
		return &Email{Address: address}, nil
	}
}
