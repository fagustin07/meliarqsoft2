package model

type CreateUUIDError struct{}

func (c CreateUUIDError) Error() string {
	return "failed generating a new id"
}

type InvalidEmailError struct{}

func (i InvalidEmailError) Error() string {
	return "invalid email address"
}
