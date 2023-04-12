package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
)

// Delete
// @Summary Delete a seller
// @Description Delete seller from a seller
// @Produce json
// @Tags Sellers
// @Param 	id 	path  string true "ID from seller to delete"
// @Success 204
// @Failure 404
// @Failure 400
// @Router /sellers/{id} [DELETE]
func (handler SellerGinHandler) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.Status(http.StatusBadRequest)
		log.Println(err.Error())
		return
	}

	err = handler.service.Delete(id)

	if err != nil {
		c.Status(http.StatusNotFound)
		log.Println(err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}
