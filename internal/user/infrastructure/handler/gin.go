package handler

import (
	"meliarqsoft2/internal/user/domain/ports"
)

type UserGinHandler struct {
	service ports.IUserService
}

func NewUserGinHandler(service ports.IUserService) *UserGinHandler {
	return &UserGinHandler{service: service}
}
