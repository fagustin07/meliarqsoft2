package handler

import (
	"meliarqsoft2/internal/user/domain/ports"
)

type UserGinHandler struct {
	manager ports.IUserManager
}

func NewUserGinHandler(manager ports.IUserManager) *UserGinHandler {
	return &UserGinHandler{manager: manager}
}
