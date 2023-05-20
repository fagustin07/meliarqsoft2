package handler

import (
	"github.com/gin-gonic/gin"
	"meliarqsoft2/internal/domain/model"
	"meliarqsoft2/pkg/exceptions/application"
	"net/http"
)

type GinFilterProduct struct {
	productService model.IProductService
}

func NewGinFilterProduct(productService model.IProductService) *GinFilterProduct {
	return &GinFilterProduct{productService: productService}
}

// Execute Filter
// @Summary Filter products
// @Description Filter products that contains given name string and category string
// @Accept json
// @Produce json
// @Tags Products
// @Param min_price query number  false "starting price" minimum(0)
// @Param max_price query number  false "limit price"    minimum(0)
// @Success 200 {object} []model.Product
// @Failure 400
// @Failure 500
// @Router /products/prices [GET]
func (handler GinFilterProduct) Execute(c *gin.Context) {
	var filterQuery model.FilterProductQuery
	if err := c.ShouldBindQuery(&filterQuery); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := handler.productService.Filter(filterQuery.GetMin(), filterQuery.GetMax())
	if err != nil {
		application.MeliGinHandlerError{}.Execute(err, c)
		return
	}

	c.JSON(http.StatusOK, resp)
}
