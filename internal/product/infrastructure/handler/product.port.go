package handler

import (
	"github.com/gin-gonic/gin"
	"meliarqsoft2/internal/product/domain"
)

type IProductHandler interface {
	Create(c *gin.Context) error
	Update(c *gin.Context) error
	Delete(c *gin.Context) error
	Find(c *gin.Context) domain.Product
	Filter(c *gin.Context) domain.Product
}
