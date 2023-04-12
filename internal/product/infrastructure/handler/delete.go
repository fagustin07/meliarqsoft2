package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
)

// Delete
// @Summary Delete a product
// @Description Delete product from a seller
// @Produce json
// @Tags Products
// @Param 	id 	path  string true "ID from product to delete"
// @Success 204
// @Failure 404
// @Failure 400
// @Router /products/{id} [DELETE]
func (p ProductGinHandler) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.Status(http.StatusBadRequest)
		log.Println(err.Error())
		return
	}

	err = p.productService.Delete(id)

	if err != nil {
		c.Status(http.StatusNotFound)
		log.Println(err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}
