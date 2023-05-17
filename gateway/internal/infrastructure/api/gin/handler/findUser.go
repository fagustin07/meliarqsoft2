package handler

import (
	"github.com/gin-gonic/gin"
	"meliarqsoft2/internal/domain/model"
	"meliarqsoft2/pkg/exceptions/application"
	"net/http"
)

type GinFindUser struct {
	UserService model.IUserService
}

func NewGinFindUser(userService model.IUserService) *GinFindUser {
	return &GinFindUser{UserService: userService}
}

// Execute Find User
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
func (handler GinFindUser) Execute(c *gin.Context) {
	emailPattern := c.Query("email")

	resp, err := handler.UserService.Find(emailPattern)
	if err != nil {
		application.MeliGinHandlerError{}.Execute(err, c)
		return
	}

	c.JSON(http.StatusOK, resp)
}
