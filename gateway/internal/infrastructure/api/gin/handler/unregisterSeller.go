package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/model"
	"meliarqsoft2/pkg/exceptions/application"
	"net/http"
)

type GinUnregisterSeller struct {
	sellerService model.ISellerService
}

func NewGinUnregisterSeller(sellerService model.ISellerService) *GinUnregisterSeller {
	return &GinUnregisterSeller{sellerService: sellerService}
}

// Execute Unregister a seller
// @Summary Unregister a seller
// @Description Unregister seller from a seller
// @Produce json
// @Tags Sellers
// @Param 	id 	path  string true "ID from seller to unregister"
// @Success 204
// @Failure 400
// @Failure 404
// @Failure 500
// @Failure 503
// @Router /sellers/{id} [DELETE]
func (handler GinUnregisterSeller) Execute(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = handler.sellerService.Delete(id)

	if err != nil {
		application.MeliGinHandlerError{}.Execute(err, c)
		return
	}

	c.Status(http.StatusNoContent)
}
