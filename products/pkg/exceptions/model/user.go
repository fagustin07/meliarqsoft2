package model

type UserNotFoundError struct {
	Id string
}

func (u UserNotFoundError) Error() string {
	return "user " + u.Id + " not found"
}

type UserAlreadyExistError struct{}

func (u UserAlreadyExistError) Error() string {
	return "email given already taken"
}
