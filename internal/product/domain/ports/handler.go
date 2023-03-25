package ports

import (
	"github.com/gin-gonic/gin"
)

type IProductHandler interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	Find(c *gin.Context)
	Filter(c *gin.Context)
}
