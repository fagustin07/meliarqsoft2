package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"meliarqsoft2/internal/application/command/action"
	dto2 "meliarqsoft2/internal/infrastructure/api/dto"
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

	id, err := handler.RegisterUserEvent.Execute(userDTO.Name, userDTO.Surname, userDTO.Email)
	if err != nil {
		c.Status(http.StatusBadRequest)
		log.Print(err)
		return
	}
	c.JSON(http.StatusCreated, dto2.UserID{ID: id})
}
