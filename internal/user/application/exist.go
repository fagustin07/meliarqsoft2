package application

import "github.com/google/uuid"

func (service UserService) Exist(idUser uuid.UUID) error {
	return service.repository.Exist(idUser)
}
