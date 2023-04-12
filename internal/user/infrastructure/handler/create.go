package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"meliarqsoft2/internal/user/infrastructure/dto"
	"net/http"
)

// Create
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
func (handler UserGinHandler) Create(c *gin.Context) {
	var userDTO dto.CreateUserRequest
	if err := c.BindJSON(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := handler.service.Create(userDTO.Name, userDTO.Surname, userDTO.Email)
	if err != nil {
		c.Status(http.StatusBadRequest)
		log.Print(err)
		return
	}
	c.JSON(http.StatusCreated, dto.UserID{ID: id})
}
