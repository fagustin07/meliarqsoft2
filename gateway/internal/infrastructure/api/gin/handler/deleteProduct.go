package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/model"
	"meliarqsoft2/pkg/exceptions/application"
	"net/http"
)

type GinDeleteProduct struct {
	productService model.IProductService
}

func NewGinDeleteProduct(productService model.IProductService) *GinDeleteProduct {
	return &GinDeleteProduct{productService: productService}
}

// Execute Delete
// @Summary Delete a product
// @Description Delete product
// @Produce json
// @Tags Products
// @Param 	id 	path  string true "ID from product to delete"
// @Success 204
// @Failure 400
// @Failure 404
// @Router /products/{id} [DELETE]
func (handler GinDeleteProduct) Execute(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = handler.productService.Delete(id)

	if err != nil {
		application.MeliGinHandlerError{}.Execute(err, c)
		return
	}

	c.Status(http.StatusNoContent)
}
