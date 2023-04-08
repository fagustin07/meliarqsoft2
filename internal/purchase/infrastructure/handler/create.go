package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"meliarqsoft2/internal/purchase/infrastructure/dto"
	"net/http"
)

// Create
// @Summary Make a purchase.
// @Description Make User buy Product by ids.
// @Accept json
// @Produce json
// @Tags Purchases
// @Param Body body dto.CreatePurchaseRequest true "Register"
// @Success 201 {object} []dto.PurchaseDTO
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /products/purchases [POST]
func (handler PurchaseGinHandler) Create(c *gin.Context) {
	var purchaseDTO dto.CreatePurchaseRequest
	if err := c.BindJSON(&purchaseDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	purchase, err := handler.manager.Create(purchaseDTO.IDProduct, purchaseDTO.IDUser, purchaseDTO.Units)
	if err != nil {
		c.Status(http.StatusBadRequest)
		log.Print(err)
		return
	}

	c.JSON(http.StatusCreated, dto.MapPurchaseToJSON(purchase))
}
