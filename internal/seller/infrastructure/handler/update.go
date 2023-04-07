package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"meliarqsoft2/internal/seller/infrastructure/dto"
	"net/http"
)

// Update
// @Summary Update a seller
// @Description Update seller from a seller
// @Accept json
// @Produce json
// @Tags Sellers
// @Param Body body dto.UpdateSellerRequest true "Register"
// @Param 	id 	path  string true "ID from seller to update"
// @Success 204
// @Failure 400
// @Failure 500
// @Router /sellers/{id} [PUT]
func (handler SellerGinHandler) Update(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		log.Println(err.Error())
		return
	}

	var dataToUpdate dto.UpdateSellerRequest
	if err := c.BindJSON(&dataToUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = handler.manager.Update(id, dataToUpdate.BusinessName, dataToUpdate.Email)

	if err != nil {
		c.Status(http.StatusBadRequest)
		log.Print(err)
		return
	}

	c.Status(http.StatusNoContent)
}
