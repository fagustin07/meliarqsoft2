package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"meliarqsoft2/internal/purchase/infrastructure/dto"
	"net/http"
)

// Find
// @Summary Find purchases.
// @Description Find purchase in a range of dates.
// @Accept json
// @Produce json
// @Tags Purchases
// @Param Body body dto.FindPurchaseRequest true "Register"
// @Success 201 {object} []dto.PurchaseDTO
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /products/purchases/prices [POST]
func (handler PurchaseGinHandler) Find(c *gin.Context) {
	var findFilter dto.FindPurchaseRequest
	if err := c.BindJSON(&findFilter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := handler.manager.Find(findFilter.Limit, findFilter.Start)
	if err != nil {
		c.Status(http.StatusBadRequest)
		log.Print(err)
		return
	}

	var res []dto.PurchaseDTO
	for _, elem := range resp {
		res = append(res, dto.MapPurchaseToJSON(elem))
	}

	c.JSON(http.StatusCreated, res)
}
