package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
)

// Delete
// @Summary Delete a user
// @Description Delete user from a user
// @Produce json
// @Tags Users
// @Param 	id 	path  string true "ID from user to delete"
// @Success 204
// @Failure 400
// @Failure 404
// @Failure 400
// @Router /users/{id} [DELETE]
func (handler UserGinHandler) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.Status(http.StatusBadRequest)
		log.Println(err.Error())
		return
	}

	err = handler.service.Delete(id)

	if err != nil {
		c.Status(http.StatusNotFound)
		log.Println(err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}
