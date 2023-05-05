package handler

import (
	"github.com/gin-gonic/gin"
	"meliarqsoft2/internal/domain/application/query"
	"meliarqsoft2/internal/infrastructure/api/dto"
	"meliarqsoft2/pkg/exceptions/application"
	"net/http"
)

type GinFindProduct struct {
	FindProductEvent *query.FindProductEvent
}

func NewGinFindProduct(findProductEvent *query.FindProductEvent) *GinFindProduct {
	return &GinFindProduct{FindProductEvent: findProductEvent}
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

	resp, err := handler.FindProductEvent.Execute(name, category)
	if err != nil {
		application.MeliGinHandlerError{}.Execute(err, c)
		return
	}

	var res []dto.ProductDTO
	for _, elem := range resp {
		res = append(res, dto.MapProductProductToJSON(elem))
	}
	c.JSON(http.StatusOK, res)
}
