package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/application/action"
	"meliarqsoft2/pkg/exceptions/application"
	"net/http"
)

type GinDeleteFromProduct struct {
	deleteFromProductEvent *action.UndoPurchasesByProductEvent
}

func NewGinDeleteFromProduct(deleteFromProductEvent *action.UndoPurchasesByProductEvent) *GinDeleteFromProduct {
	return &GinDeleteFromProduct{deleteFromProductEvent: deleteFromProductEvent}
}

// Execute Delete
// @Summary Delete all purchases from product
// @Description Delete all purchases from product
// @Produce json
// @Tags Purchases
// @Param 	IDs body ProductsIDs true "Register"
// @Success 204
// @Failure 409
// @Failure 400
// @Failure 500
// @Router /purchases [DELETE]
func (handler GinDeleteFromProduct) Execute(c *gin.Context) {
	var productsIDS ProductsIDs
	if err := c.BindJSON(&productsIDS); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := handler.deleteFromProductEvent.Execute(productsIDS.IDs)

	if err != nil {
		application.MeliGinHandlerError{}.Execute(err, c)
		return
	}

	c.Status(http.StatusNoContent)
}

type ProductsIDs struct {
	IDs []uuid.UUID `json:"ids"`
}
