package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"meliarqsoft2/internal/domain/application/command/action"
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
		c.Status(http.StatusBadRequest)
		log.Println(err.Error())
		return
	}

	err = handler.deleteProductEvent.Execute(id)

	if err != nil {
		c.Status(http.StatusNotFound)
		log.Println(err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}
