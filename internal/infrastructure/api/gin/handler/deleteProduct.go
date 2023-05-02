package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/application/action"
	"meliarqsoft2/pkg/exceptions/application"
	"net/http"
)

type GinDeleteProduct struct {
	deleteProductEvent *action.DeleteProductEvent
}

func NewGinDeleteProduct(deleteProductEvent *action.DeleteProductEvent) *GinDeleteProduct {
	return &GinDeleteProduct{deleteProductEvent: deleteProductEvent}
}

// Execute Delete
// @Summary Delete a product
// @Description Delete product from a seller
// @Produce json
// @Tags Products
// @Param 	id 	path  string true "ID from product to delete"
// @Success 204
// @Failure 404
// @Failure 400
// @Router /products/{id} [DELETE]
func (handler GinDeleteProduct) Execute(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = handler.deleteProductEvent.Execute(id)

	if err != nil {
		application.MeliGinHandlerError{}.Execute(err, c)
		return
	}

	c.Status(http.StatusNoContent)
}
