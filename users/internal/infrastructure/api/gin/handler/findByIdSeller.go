package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/application/query"
	"meliarqsoft2/internal/infrastructure/api/dto"
	"meliarqsoft2/pkg/exceptions/application"
	"net/http"
)

type GinFindByIdSeller struct {
	FindByIdSellerEvent *query.FindSellerByIdEvent
}

func NewGinFindByIdSeller(findByIdSellerEvent *query.FindSellerByIdEvent) *GinFindByIdSeller {
	return &GinFindByIdSeller{FindByIdSellerEvent: findByIdSellerEvent}
}

// Execute FindById seller
// @Summary FindById a seller
// @Description FindById seller
// @Accept json
// @Produce json
// @Tags Sellers
// @Param 	id 	path  string true "ID from seller"
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Failure 503
// @Router /sellers/{id} [GET]
func (handler GinFindByIdSeller) Execute(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	seller, err := handler.FindByIdSellerEvent.Execute(id)

	if err != nil {
		application.MeliGinHandlerError{}.Execute(err, c)
		return
	}

	c.JSON(http.StatusOK, dto.MapSellerToJSON(seller))
}
