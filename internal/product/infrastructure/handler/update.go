package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"meliarqsoft2/internal/product/infrastructure/dto"
	"net/http"
)

// Update
// @Summary Update a product
// @Description Update product from a seller
// @Accept json
// @Produce json
// @Tags Products
// @Param Body body dto.UpdateProductRequest true "Register"
// @Param 	id 	path  string true "ID from product to update"
// @Success 204
// @Failure 400
// @Failure 500
// @Router /products/{id} [PUT]
func (p ProductGinHandler) Update(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		log.Println(err.Error())
		return
	}

	var dataToUpdate dto.UpdateProductRequest
	if err := c.BindJSON(&dataToUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = p.productService.Update(id, dataToUpdate.Name, dataToUpdate.Description, dataToUpdate.Category, dataToUpdate.Price, dataToUpdate.Stock)
	if err != nil {
		c.Status(http.StatusBadRequest)
		log.Print(err)
		return
	}

	c.Status(http.StatusNoContent)
}
