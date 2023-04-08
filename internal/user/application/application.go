package application

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/user/domain"
	"meliarqsoft2/internal/user/domain/ports"
)

type UserManager struct {
	repository ports.IUserRepository
}

func NewUserManager(repo ports.IUserRepository) *UserManager {
	return &UserManager{repository: repo}
}
func (manager UserManager) Create(name string, surname string, email string) (uuid.UUID, error) {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return newUUID, err
	}

	newUser := domain.NewUser(newUUID, name, surname, email)
	err = manager.repository.Create(newUser)

	if err != nil {
		return newUUID, err
	}

	return newUUID, nil
}

func (manager UserManager) Update(ID uuid.UUID, name string, surname string, email string) error {
	return manager.repository.Update(ID, name, surname, email)
}

func (manager UserManager) Delete(ID uuid.UUID) error {
	return manager.repository.Delete(ID)
}

func (manager UserManager) Find(ID uuid.UUID) (*domain.User, error) {
	return manager.repository.Find(ID)
}
