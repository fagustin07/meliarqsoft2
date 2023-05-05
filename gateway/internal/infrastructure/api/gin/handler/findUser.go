package handler

import (
	"github.com/gin-gonic/gin"
	"meliarqsoft2/internal/domain/application/query"
	"meliarqsoft2/internal/infrastructure/api/dto"
	"meliarqsoft2/pkg/exceptions/application"
	"net/http"
)

type GinFindUser struct {
	FindUserEvent *query.FindUserEvent
}

func NewGinFindUser(findUserEvent *query.FindUserEvent) *GinFindUser {
	return &GinFindUser{FindUserEvent: findUserEvent}
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

	resp, err := handler.FindUserEvent.Execute(emailPattern)
	if err != nil {
		application.MeliGinHandlerError{}.Execute(err, c)
		return
	}

	var res []dto.UserDTO
	for _, elem := range resp {
		res = append(res, dto.MapUserToJSON(elem))
	}

	c.JSON(http.StatusOK, res)
}
