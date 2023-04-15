package ports

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain"
)

type IUserService interface {
	Create(name string, surname string, email string) (uuid.UUID, error)
	Update(ID uuid.UUID, name string, surname string, email string) error
	Delete(ID uuid.UUID) error
	Find(emailPattern string) ([]*domain.User, error)
	Exist(idUser uuid.UUID) error
}
