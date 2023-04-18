package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"meliarqsoft2/internal/application/command/action"
	dto2 "meliarqsoft2/internal/infrastructure/api/dto"
	"net/http"
)

type GinCreateProduct struct {
	CreateProductEvent *action.CreateProductEvent
}

func NewGinCreateProduct(createProductEvent *action.CreateProductEvent) *GinCreateProduct {
	return &GinCreateProduct{CreateProductEvent: createProductEvent}
}

// Execute Create
// @Summary Create a product
// @Description Create product from a seller
// @Accept json
// @Produce json
// @Tags Products
// @Param Body body dto.CreateProductRequest true "Create a product"
// @Success 200 {object} dto.ProductID
// @Failure 400
// @Failure 500
// @Router /products [post]
func (handler GinCreateProduct) Execute(c *gin.Context) {
	var productDTO dto2.CreateProductRequest
	if err := c.BindJSON(&productDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := productDTO.MapToModel()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdProductID, err := handler.CreateProductEvent.Execute(product)
	if err != nil {
		c.Status(http.StatusBadRequest)
		log.Print(err)
		return
	}
	c.JSON(http.StatusCreated, dto2.ProductID{ID: createdProductID})
}
