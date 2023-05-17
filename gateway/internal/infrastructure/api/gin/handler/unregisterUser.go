package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/model"
	"meliarqsoft2/pkg/exceptions/application"
	"net/http"
)

type GinUnregisterUser struct {
	UserService model.IUserService
}

func NewGinUnregisterUser(userService model.IUserService) *GinUnregisterUser {
	return &GinUnregisterUser{UserService: userService}
}

// Execute Unregister User
// @Summary Unregister a user
// @Description Unregister user from a user
// @Produce json
// @Tags Users
// @Param 	id 	path  string true "ID from user to unregister"
// @Success 204
// @Failure 400
// @Failure 404
// @Failure 400
// @Router /users/{id} [DELETE]
func (handler GinUnregisterUser) Execute(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = handler.UserService.Delete(id)

	if err != nil {
		application.MeliGinHandlerError{}.Execute(err, c)
		return
	}

	c.Status(http.StatusNoContent)
}
