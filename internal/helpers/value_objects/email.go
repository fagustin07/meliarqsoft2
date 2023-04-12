package value_objects

import (
	"errors"
	"net/mail"
)

type Email struct {
	Address string
}

func NewEmail(address string) (*Email, error) {
	addr, err := mail.ParseAddress(address)

	if err != nil {
		return nil, errors.New("invalid email address")
	} else {
		return &Email{Address: addr.String()}, nil
	}
}
