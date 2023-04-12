package application

import (
	"meliarqsoft2/internal/user/domain/ports"
)

type UserService struct {
	repository ports.IUserRepository
}

func NewUserService(repo ports.IUserRepository) *UserService {
	return &UserService{repository: repo}
}
