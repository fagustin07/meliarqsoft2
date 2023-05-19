package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/model"
	"meliarqsoft2/pkg/exceptions/application"
	"net/http"
)

type GinFindPurchases struct {
	findPurchaseService model.IFindPurchaseService
}

func NewGinFindPurchases(findPurchaseService model.IFindPurchaseService) *GinFindPurchases {
	return &GinFindPurchases{findPurchaseService: findPurchaseService}
}

// Execute Find
// @Summary Find purchases from product.
// @Description Find purchases from product
// @Produce json
// @Tags Purchases
// @Param 	id 	path  string true "ID from product"
// @Success 200 {object} []model.Purchase
// @Failure 404
// @Failure 400
// @Router /products/{id}/purchases [GET]
func (handler GinFindPurchases) Execute(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := handler.findPurchaseService.Execute(id)
	if err != nil {
		application.MeliGinHandlerError{}.Execute(err, c)
		return
	}

	c.JSON(http.StatusOK, resp)
}
