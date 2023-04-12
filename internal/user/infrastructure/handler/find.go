package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"meliarqsoft2/internal/user/infrastructure/dto"
	"net/http"
)

// Find
// @Summary Find user
// @Description Find user with given id
// @Accept json
// @Produce json
// @Tags Users
// @Param email	query string  false "find name"
// @Success 200 {object} []dto.UserDTO
// @Failure 400
// @Failure 500
// @Router /users [GET]
func (handler UserGinHandler) Find(c *gin.Context) {
	name := c.Query("email")

	resp, err := handler.service.Find(name)
	if err != nil {
		c.Status(http.StatusNotFound)
		log.Println(err.Error())
		return
	}

	var res []dto.UserDTO
	for _, elem := range resp {
		res = append(res, dto.MapUserToJSON(elem))
	}

	c.JSON(http.StatusOK, res)
}
