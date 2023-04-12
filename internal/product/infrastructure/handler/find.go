package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"meliarqsoft2/internal/product/infrastructure/dto"
	"net/http"
)

// Find
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
func (p ProductGinHandler) Find(c *gin.Context) {
	name := c.Query("name")
	category := c.Query("category")

	resp, err := p.productService.Find(name, category)
	if err != nil {
		c.Status(http.StatusBadRequest)
		log.Print(err)
		return
	}

	var res []dto.ProductDTO
	for _, elem := range resp {
		res = append(res, dto.MapProductProductToJSON(elem))
	}
	c.JSON(http.StatusOK, res)
}
