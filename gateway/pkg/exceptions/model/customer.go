package model

type CustomerNotFoundError struct {
	Id string
}

func (u CustomerNotFoundError) Error() string {
	return "customer " + u.Id + " not found"
}

type CustomerAlreadyExistError struct{}

func (u CustomerAlreadyExistError) Error() string {
	return "email given already taken"
}
