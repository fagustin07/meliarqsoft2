package ports

import (
	"github.com/gin-gonic/gin"
	"meliarqsoft2/internal/core/domain"
)

type IProductService interface {
	Create(name string, description string, price float32, stock int, idSeller int) error
	Update(ID int, name string, description string, price float32, stock int, idSeller int) error
	Delete(ID int) error
	Find(name string, category string) domain.Product
	Filter(minPrice float32, maxPrice float32) domain.Product
}

type IProductRepository interface {
	Create(ID int, name string, description string, price float32, stock int, idSeller int) error
	Update(ID int, name string, description string, price float32, stock int, idSeller int) error
	Delete(ID int) error
	Find(name string, category string) domain.Product
	Filter(minPrice float32, maxPrice float32) domain.Product
}

type IProductHandler interface {
	Create(c *gin.Context) error
	Update(c *gin.Context) error
	Delete(c *gin.Context) error
	Find(c *gin.Context) domain.Product
	Filter(c *gin.Context) domain.Product
}
