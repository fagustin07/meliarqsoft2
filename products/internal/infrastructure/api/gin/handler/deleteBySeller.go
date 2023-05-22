package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/application/action"
	"meliarqsoft2/pkg/exceptions/application"
	"net/http"
)

type GinDeleteProductsBySeller struct {
	DeleteProductsBySeller *action.DeleteProductsBySeller
}

func NewGinDeleteProductsBySeller(deleteByProducts *action.DeleteProductsBySeller) *GinDeleteProductsBySeller {
	return &GinDeleteProductsBySeller{DeleteProductsBySeller: deleteByProducts}
}

// Execute Delete
// @Summary Delete products from a seller
// @Description Delete products from a seller
// @Tags 	Products
// @Param 	id 	path  string true "ID from seller"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /products/sellers/{id} [DELETE]
func (handler GinDeleteProductsBySeller) Execute(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ids, err := handler.DeleteProductsBySeller.Execute(id)
	if err != nil {
		application.MeliGinHandlerError{}.Execute(err, c)
		return
	}

	c.JSON(http.StatusOK, ProductsIDs{IDs: ids})
}
