package handler

import (
	"github.com/gin-gonic/gin"
	"meliarqsoft2/internal/domain/application/action"
	dto2 "meliarqsoft2/internal/infrastructure/api/dto"
	"meliarqsoft2/pkg/exceptions/application"
	"net/http"
)

type GinRegisterUser struct {
	RegisterUserEvent *action.RegisterUserEvent
}

func NewGinUserRegister(registerUserEvent *action.RegisterUserEvent) *GinRegisterUser {
	return &GinRegisterUser{RegisterUserEvent: registerUserEvent}
}

// Execute Register User
// @Summary Create a user
// @Description Create user
// @Accept json
// @Produce json
// @Tags Users
// @Param Body body dto.CreateUserRequest true "Register"
// @Success 201 {object} dto.UserID
// @Failure 400
// @Failure 500
// @Router /users [POST]
func (handler GinRegisterUser) Execute(c *gin.Context) {
	var userDTO dto2.CreateUserRequest
	if err := c.BindJSON(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	toModel, err := userDTO.MapToModel()
	if err != nil {
		application.MeliGinHandlerError{}.Execute(err, c)
		return
	}

	id, err := handler.RegisterUserEvent.Execute(&toModel)
	if err != nil {
		application.MeliGinHandlerError{}.Execute(err, c)
		return
	}
	c.JSON(http.StatusCreated, dto2.UserID{ID: id})
}
