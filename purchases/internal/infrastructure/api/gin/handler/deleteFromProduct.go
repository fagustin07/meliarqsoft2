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
// @Param 	id 	path  string true "ID from product to delete purchases"
// @Success 204
// @Failure 404
// @Failure 400
// @Router /products/{id} [DELETE]
func (handler GinDeleteFromProduct) Execute(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = handler.deleteFromProductEvent.Execute(id)

	if err != nil {
		application.MeliGinHandlerError{}.Execute(err, c)
		return
	}

	c.Status(http.StatusNoContent)
}
