package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"meliarqsoft2/internal/purchase/infrastructure/dto"
	"net/http"
)

// Find
// @Summary Find purchases from product.
// @Description Find purchases from product
// @Produce json
// @Tags Purchases
// @Param 	id 	path  string true "ID from product"
// @Success 200
// @Failure 404
// @Failure 400
// @Router /products/{id}/purchases [GET]
func (handler PurchaseGinHandler) Find(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	log.Println(id)
	if err != nil {
		c.Status(http.StatusBadRequest)
		log.Println(err.Error())
		return
	}
	resp, err := handler.purchaseService.Find(id)
	if err != nil {
		c.Status(http.StatusBadRequest)
		log.Print(err)
		return
	}

	var res []dto.PurchaseDTO
	for _, elem := range resp {
		res = append(res, dto.MapPurchaseToJSON(elem))
	}

	c.JSON(http.StatusOK, res)
}
