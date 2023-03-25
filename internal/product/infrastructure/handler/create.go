package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"meliarqsoft2/internal/product/infrastructure/dto"
	"net/http"
)

func (p ProductGinHandler) Create(c *gin.Context) {
	var productDTO dto.CreateProductRequest
	if err := c.BindJSON(&productDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println("HERE", productDTO.Name, productDTO.IDSeller)

	createdProduct, err := p.app.Create(productDTO.Name, productDTO.Description, productDTO.Category, productDTO.Price, productDTO.Stock, productDTO.IDSeller)
	if err != nil {
		c.Status(http.StatusBadRequest)
		log.Fatalln(err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"product": dto.MapCreatedProductProductToJSON(createdProduct)})
}
