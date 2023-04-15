package application

import (
	"meliarqsoft2/internal/domain"
)

func (service UserService) Find(emailPattern string) ([]*domain.User, error) {
	return service.repository.Find(emailPattern)
}
