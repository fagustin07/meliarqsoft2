package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"meliarqsoft2/internal/domain/model"
	"meliarqsoft2/pkg/exceptions/application"
	"net/http"
)

type GinMakePurchase struct {
	makePurchaseService model.IMakePurchaseService
}

func NewGinMakePurchase(makePurchaseService model.IMakePurchaseService) *GinMakePurchase {
	return &GinMakePurchase{makePurchaseService: makePurchaseService}
}

// Execute Create
// @Summary Make a purchase.
// @Description Make User buy Product by ids.
// @Accept json
// @Produce json
// @Tags Purchases
// @Param 	Purhcase body model.CreatePurchaseRequest true "Register"
// @Success 201 {object} model.Purchase
// @Failure 400
// @Failure 404
// @Failure 406
// @Failure 500
// @Failure 503
// @Router /purchases/products [POST]
func (handler GinMakePurchase) Execute(c *gin.Context) {
	var purchaseReq model.CreatePurchaseRequest

	if err := c.BindJSON(&purchaseReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println(purchaseReq.Units)
	log.Println(purchaseReq.IDUser)
	log.Println(purchaseReq.IDProduct)
	resp, err := handler.makePurchaseService.Execute(purchaseReq)
	if err != nil {
		application.MeliGinHandlerError{}.Execute(err, c)
		return
	}
	c.JSON(http.StatusCreated, resp)
}
