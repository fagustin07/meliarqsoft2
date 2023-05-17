package handler

import (
	"github.com/gin-gonic/gin"
	"meliarqsoft2/internal/domain/model"
	"meliarqsoft2/pkg/exceptions/application"
	"net/http"
)

type GinFindProduct struct {
	productService model.IProductService
}

func NewGinFindProduct(productService model.IProductService) *GinFindProduct {
	return &GinFindProduct{productService: productService}
}

// Execute Find product
// @Summary Find products
// @Description Find products that contains given name string and category string
// @Accept json
// @Produce json
// @Tags Products
// @Param name 	 	query string  false "find name"
// @Param category  query string  false "find category"
// @Success 200 {object} []dto.ProductDTO
// @Failure 400
// @Failure 500
// @Router /products [get]
func (handler GinFindProduct) Execute(c *gin.Context) {
	name := c.Query("name")
	category := c.Query("category")

	resp, err := handler.productService.Find(name, category)
	if err != nil {
		application.MeliGinHandlerError{}.Execute(err, c)
		return
	}

	c.JSON(http.StatusOK, resp)
}
