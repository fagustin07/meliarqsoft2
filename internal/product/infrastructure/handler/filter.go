package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"meliarqsoft2/internal/product/infrastructure/dto"
	"net/http"
	"strconv"
)

// Filter
// @Summary Filter products
// @Description Filter products that contains given name string and category string
// @Accept json
// @Produce json
// @Tags Products
// @Param min_price query string  true "starting price"
// @Param max_price query string  true "limit price"
// @Success 200 {object} []dto.ProductDTO
// @Failure 400
// @Failure 500
// @Router /products/prices [GET]
func (p ProductGinHandler) Filter(c *gin.Context) {
	var (
		min, max float32
		err      error
	)
	if min, err = findQueryParam("min_price", c); err != nil {
		c.Status(http.StatusBadRequest)
		log.Print(err)
		return
	}

	if max, err = findQueryParam("max_price", c); err != nil {
		c.Status(http.StatusBadRequest)
		log.Print(err)
		return
	}
	resp, err := p.app.Filter(min, max)
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

func findQueryParam(key string, c *gin.Context) (float32, error) {
	val, err := strconv.ParseFloat(c.Query(key), 32)
	if err != nil {
		return 0, err
	}
	return float32(val), nil
}
