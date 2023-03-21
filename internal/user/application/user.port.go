package application

type IUserUseCase interface {
	Create(name string, email string) error
}
