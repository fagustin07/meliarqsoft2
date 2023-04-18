package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"meliarqsoft2/internal/domain/application/action"
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
		c.Status(http.StatusBadRequest)
		log.Println(err.Error())
		return
	}

	err = handler.UnregisterSellerEvent.Execute(id)

	if err != nil {
		c.Status(http.StatusNotFound)
		log.Println(err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}
