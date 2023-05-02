package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/application/action"
	"meliarqsoft2/internal/infrastructure/api/dto"
	"meliarqsoft2/pkg/exceptions/application"
	"net/http"
)

type GinUpdateUser struct {
	UpdateUserEvent *action.UpdateUserEvent
}

func NewGinUpdateUser(unregisterUserEvent *action.UpdateUserEvent) *GinUpdateUser {
	return &GinUpdateUser{UpdateUserEvent: unregisterUserEvent}
}

// Execute Update
// @Summary Update a user
// @Description Update user from a user
// @Produce json
// @Tags Users
// @Param Body body dto.UpdateUserRequest true "Register"
// @Param 	id 	path  string true "ID from user to update"
// @Success 204
// @Failure 404
// @Failure 400
// @Router /users/{id} [PUT]
func (handler GinUpdateUser) Execute(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var dataToUpdate dto.UpdateUserRequest
	if err := c.BindJSON(&dataToUpdate); err != nil {
		application.MeliGinHandlerError{}.Execute(err, c)
		return
	}

	err = handler.UpdateUserEvent.Execute(id, dataToUpdate.Name, dataToUpdate.Surname, dataToUpdate.Email)

	if err != nil {
		application.MeliGinHandlerError{}.Execute(err, c)
		return
	}

	c.Status(http.StatusNoContent)
}
