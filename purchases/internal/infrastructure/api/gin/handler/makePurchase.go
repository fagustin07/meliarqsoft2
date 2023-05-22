package handler

import (
	"github.com/gin-gonic/gin"
	"meliarqsoft2/internal/domain/application/action"
	dto2 "meliarqsoft2/internal/infrastructure/api/dto"
	"meliarqsoft2/pkg/exceptions/application"
	"net/http"
)

type GinMakePurchase struct {
	MakePurchaseEvent     *action.MakePurchaseEvent
	SendNotificationEvent *action.SendNotificationEvent
}

func NewGinMakePurchase(makePurchaseEvent *action.MakePurchaseEvent, sendNotificationEvent *action.SendNotificationEvent) *GinMakePurchase {
	return &GinMakePurchase{
		MakePurchaseEvent:     makePurchaseEvent,
		SendNotificationEvent: sendNotificationEvent,
	}
}

// Execute Create
// @Summary Make a purchase.
// @Description Make User buy Product by ids.
// @Accept json
// @Produce json
// @Tags Purchases
// @Param 	Purchase body dto.CreatePurchaseRequest true "Register"
// @Success 201 {object} []dto.PurchaseDTO
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /purchases [POST]
func (handler GinMakePurchase) Execute(c *gin.Context) {
	var purchaseDTO dto2.CreatePurchaseRequest
	if err := c.BindJSON(&purchaseDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	purchase, sellerNotification, buyerNotification, err := purchaseDTO.MapToModel()
	if err != nil {
		application.MeliGinHandlerError{}.Execute(err, c)
		return
	}

	_, err = handler.MakePurchaseEvent.Execute(&purchase)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Enviar notificaciones

	errSeller := handler.SendNotificationEvent.Execute(&sellerNotification)

	if errSeller != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	errBuyer := handler.SendNotificationEvent.Execute(&buyerNotification)

	if errBuyer != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto2.MapPurchaseToJSON(&purchase))
}
