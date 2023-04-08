package ports

import "github.com/gin-gonic/gin"

type IPurchaseHandler interface {
	Create(c *gin.Context)
	Find(c *gin.Context)
}
