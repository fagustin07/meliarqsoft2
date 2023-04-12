package application

import "github.com/google/uuid"

func (service UserService) Delete(ID uuid.UUID) error {
	return service.repository.Delete(ID)
}
