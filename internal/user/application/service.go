package application

import (
	"meliarqsoft2/internal/domain"
)

type UserService struct {
	repository domain.IUserRepository
}

func NewUserService(repo domain.IUserRepository) *UserService {
	return &UserService{repository: repo}
}
