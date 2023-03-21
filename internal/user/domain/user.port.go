package domain

type IUserRepository interface {
	Create(ID int, name string, email string) error
}
