package ports

import "github.com/gin-gonic/gin"

type ISellerHandler interface {
	Create(c *gin.Context) error
}
