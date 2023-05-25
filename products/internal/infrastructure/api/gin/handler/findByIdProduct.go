package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/application/query"
	"meliarqsoft2/internal/infrastructure/api/dto"
	"meliarqsoft2/pkg/exceptions/application"
	"net/http"
)

type GinFindProductById struct {
	FindProductById *query.FindProductById
}

func NewGinFindProdById(findProdEvent *query.FindProductById) *GinFindProductById {
	return &GinFindProductById{FindProductById: findProdEvent}
}

// Execute Find
// @Summary Find a product
// @Description Find product
// @Accept 	json
// @Produce json
// @Tags 	Products
// @Param 	id 	path  string true "ID from product"
// @Success 200 {object} dto.ProductDTO
// @Failure 400
// @Failure 500
// @Router /products/{id} [GET]
func (handler GinFindProductById) Execute(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	prod, err := handler.FindProductById.Execute(id)
	if err != nil {
		application.MeliGinHandlerError{}.Execute(err, c)
		return
	}

	c.JSON(http.StatusOK, dto.MapProductProductToJSON(*prod))
}
