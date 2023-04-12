package helpers

import (
	"errors"
	"net/mail"
)

type Email struct {
	address string
}

func NewEmail(address string) (Email, error) {
	addr, err := mail.ParseAddress(address)

	if err != nil {
		return Email{}, errors.New("invalid email address")
	} else {
		return Email{address: addr.String()}, nil
	}
}
