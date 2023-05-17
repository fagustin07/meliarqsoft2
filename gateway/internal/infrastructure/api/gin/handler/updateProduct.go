package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/model"
	"meliarqsoft2/pkg/exceptions/application"
	"net/http"
)

type GinUpdateProduct struct {
	productService model.IProductService
}

func NewGinUpdateProduct(productService model.IProductService) *GinUpdateProduct {
	return &GinUpdateProduct{productService: productService}
}

// Execute Update
// @Summary Update a product
// @Description Update product from a seller
// @Accept 	json
// @Produce json
// @Tags 	Products
// @Param 	Body body dto.UpdateProductRequest true "Register"
// @Param 	id 	path  string true "ID from product to update"
// @Success 204
// @Failure 400
// @Failure 500
// @Router /products/{id} [PUT]
func (handler GinUpdateProduct) Execute(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var dataToUpdate model.UpdateProductRequest
	if err := c.BindJSON(&dataToUpdate); err != nil {
		application.MeliGinHandlerError{}.Execute(err, c)
		return
	}

	err = handler.productService.Update(id, dataToUpdate)

	if err != nil {
		application.MeliGinHandlerError{}.Execute(err, c)
		return
	}

	c.Status(http.StatusNoContent)
}
