package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/application/action"
	"meliarqsoft2/internal/infrastructure/api/dto"
	"meliarqsoft2/pkg/exceptions/application"
	"net/http"
)

type GinUpdateSeller struct {
	updateSellerEvent *action.UpdateSellerEvent
}

func NewGinUpdateSeller(updateSellerEvent *action.UpdateSellerEvent) *GinUpdateSeller {
	return &GinUpdateSeller{updateSellerEvent: updateSellerEvent}
}

// Execute Update seller
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
func (handler GinUpdateSeller) Execute(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var dataToUpdate dto.UpdateSellerRequest
	if err := c.BindJSON(&dataToUpdate); err != nil {
		application.MeliGinHandlerError{}.Execute(err, c)
		return
	}

	err = handler.updateSellerEvent.Execute(id, dataToUpdate.BusinessName, dataToUpdate.Email)

	if err != nil {
		application.MeliGinHandlerError{}.Execute(err, c)
		return
	}

	c.Status(http.StatusNoContent)
}
