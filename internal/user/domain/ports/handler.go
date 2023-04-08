package ports

import (
	"github.com/gin-gonic/gin"
)

type IUserHandler interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	Find(c *gin.Context)
}
