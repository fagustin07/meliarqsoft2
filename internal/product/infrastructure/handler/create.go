package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"meliarqsoft2/internal/product/infrastructure/dto"
	"net/http"
)

// Create
// @Summary Create a product
// @Description Create product from a seller
// @Accept json
// @Produce json
// @Tags Products
// @Param Body body dto.CreateProductRequest true "Register"
// @Success 200 {object} dto.ProductID
// @Failure 400
// @Failure 500
// @Router /products [post]
func (p ProductGinHandler) Create(c *gin.Context) {
	var productDTO dto.CreateProductRequest
	if err := c.BindJSON(&productDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdProduct, err := p.app.Create(productDTO.Name, productDTO.Description, productDTO.Category, productDTO.Price, productDTO.Stock, productDTO.IDSeller)
	if err != nil {
		c.Status(http.StatusBadRequest)
		log.Print(err)
		return
	}
	c.JSON(http.StatusCreated, dto.ProductID{ID: createdProduct.ID})
}
