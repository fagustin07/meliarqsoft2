package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/application/action"
	"meliarqsoft2/pkg/exceptions/application"
	"net/http"
)

type GinUnregisterSeller struct {
	UnregisterSellerEvent *action.UnregisterSellerEvent
}

func NewGinUnregisterSeller(unregisterSellerEvent *action.UnregisterSellerEvent) *GinUnregisterSeller {
	return &GinUnregisterSeller{UnregisterSellerEvent: unregisterSellerEvent}
}

// Execute Unregister a seller
// @Summary Unregister a seller
// @Description Unregister seller from a seller
// @Produce json
// @Tags Sellers
// @Param 	id 	path  string true "ID from seller to unregister"
// @Success 204
// @Failure 404
// @Failure 400
// @Router /sellers/{id} [DELETE]
func (handler GinUnregisterSeller) Execute(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = handler.UnregisterSellerEvent.Execute(id)

	if err != nil {
		application.MeliGinHandlerError{}.Execute(err, c)
		return
	}

	c.Status(http.StatusNoContent)
}
