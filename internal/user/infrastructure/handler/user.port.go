package handler

import "github.com/gin-gonic/gin"

type IUserHandler interface {
	Create(c *gin.Context) error
}
