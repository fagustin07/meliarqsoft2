package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"meliarqsoft2/internal/seller/infrastructure/dto"
	"net/http"
)

// Find
// @Summary Find sellers
// @Description Find sellers that contains given business name
// @Accept json
// @Produce json
// @Tags Sellers
// @Param business_name	query string  false "find name"
// @Success 200 {object} []dto.SellerDTO
// @Failure 400
// @Failure 500
// @Router /sellers [get]
func (handler SellerGinHandler) Find(c *gin.Context) {
	name := c.Query("business_name")

	resp, err := handler.manager.Find(name)
	if err != nil {
		c.Status(http.StatusBadRequest)
		log.Print(err)
		return
	}

	var res []dto.SellerDTO
	for _, elem := range resp {
		res = append(res, dto.MapSellerToJSON(elem))
	}
	c.JSON(http.StatusOK, res)
}
