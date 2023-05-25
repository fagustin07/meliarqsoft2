package handler

import (
	"github.com/gin-gonic/gin"
	"meliarqsoft2/internal/domain/application/query"
	"meliarqsoft2/internal/infrastructure/api/dto"
	"meliarqsoft2/pkg/exceptions/application"
	"net/http"
)

type GinFilterProduct struct {
	FilterProductEvent *query.FilterProductEvent
}

func NewGinFilterProduct(filterProductEvent *query.FilterProductEvent) *GinFilterProduct {
	return &GinFilterProduct{FilterProductEvent: filterProductEvent}
}

// Execute Filter
// @Summary Filter products
// @Description Filter products that contains given name string and category string
// @Accept json
// @Produce json
// @Tags Products
// @Param min_price query number  false "starting price" minimum(0)
// @Param max_price query number  false "limit price"    minimum(0)
// @Success 200 {object} []dto.ProductDTO
// @Failure 400
// @Failure 500
// @Router /products/prices [GET]
func (handler GinFilterProduct) Execute(c *gin.Context) {
	var filterQuery dto.FilterProductQuery
	if err := c.ShouldBindQuery(&filterQuery); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := handler.FilterProductEvent.Execute(filterQuery.GetMin(), filterQuery.GetMax())
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
