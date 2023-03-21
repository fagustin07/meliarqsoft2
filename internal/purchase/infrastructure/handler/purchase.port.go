package handler

import (
	"github.com/gin-gonic/gin"
	"meliarqsoft2/internal/product/domain"
	domain2 "meliarqsoft2/internal/user/domain"
)

type IPurchaseHandler interface {
	Create(c *gin.Context) error
	Update(c *gin.Context) error
	Find(c *gin.Context) (domain.Product, domain2.User)
}
