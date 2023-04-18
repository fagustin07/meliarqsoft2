package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"meliarqsoft2/internal/domain/application/action"
	"net/http"
)

type GinUnregisterUser struct {
	UnregisterUserEvent *action.UnregisterUserEvent
}

func NewGinUnregisterUser(unregisterUserEvent *action.UnregisterUserEvent) *GinUnregisterUser {
	return &GinUnregisterUser{UnregisterUserEvent: unregisterUserEvent}
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
		c.Status(http.StatusBadRequest)
		log.Println(err.Error())
		return
	}

	err = handler.UnregisterUserEvent.Execute(id)

	if err != nil {
		c.Status(http.StatusNotFound)
		log.Println(err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}
