package ports

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/user/domain"
)

type IUserRepository interface {
	Create(user *domain.User) error
	Update(ID uuid.UUID, name string, surname string, email string) error
	Delete(ID uuid.UUID) error
	Find(ID uuid.UUID) (*domain.User, error)
}
