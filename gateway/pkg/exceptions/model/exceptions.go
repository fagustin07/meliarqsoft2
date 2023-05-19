package model

type CreateUUIDError struct{}

func (c CreateUUIDError) Error() string {
	return "failed generating a new id"
}

type InvalidEmailError struct{}

func (i InvalidEmailError) Error() string {
	return "invalid email address"
}

type ServiceUnavailable struct{}

func (s ServiceUnavailable) Error() string {
	return "service currently unavailable"
}

type BadRequestError struct {
	Message string `json:"error"`
}

func (s BadRequestError) Error() string {
	return s.Message
}

type NotAcceptableError struct {
	Message string `json:"error"`
}

func (s NotAcceptableError) Error() string {
	return s.Message
}
