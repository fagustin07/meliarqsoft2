package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"meliarqsoft2/internal/user/domain"
	"meliarqsoft2/internal/user/infrastructure/dto"
	"net/http"
)

// Find
// @Summary Find user
// @Description Find user with given id
// @Accept json
// @Produce json
// @Tags Users
// @Param 	id 	path  string true "ID from user to get"
// @Success 200 {object} dto.UserDTO
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /users/{id} [GET]
func (handler UserGinHandler) Find(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.Status(http.StatusBadRequest)
		log.Println(err.Error())
		return
	}

	var user *domain.User
	user, err = handler.manager.Find(id)

	if err != nil {
		c.Status(http.StatusNotFound)
		log.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, dto.MapUserToJSON(user))
}
