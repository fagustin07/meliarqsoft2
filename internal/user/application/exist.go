package application

import "github.com/google/uuid"

func (service UserService) Exist(idUser uuid.UUID) error {
	_, err := service.repository.FindById(idUser)
	return err
}
