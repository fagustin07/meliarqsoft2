package handler

import (
	"github.com/gin-gonic/gin"
	"meliarqsoft2/internal/domain/model"
	"meliarqsoft2/pkg/exceptions/application"
	"net/http"
)

type GinCreateProduct struct {
	prodService model.IProductService
}

func NewGinCreateProduct(prodService model.IProductService) *GinCreateProduct {
	return &GinCreateProduct{prodService: prodService}
}

// Execute Create
// @Summary Create a product
// @Description Create product from a seller
// @Accept json
// @Produce json
// @Tags Products
// @Param Product body model.CreateProductRequest true "Create a product"
// @Success 201 {object} model.ProductID
// @Failure 400
// @Failure 404
// @Failure 406
// @Failure 409
// @Failure 500
// @Failure 503
// @Router /products [post]
func (handler GinCreateProduct) Execute(c *gin.Context) {
	var productDTO model.CreateProductRequest
	if err := c.BindJSON(&productDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	prodId, err := handler.prodService.Create(productDTO)

	if err != nil {
		application.MeliGinHandlerError{}.Execute(err, c)
		return
	}
	c.JSON(http.StatusCreated, prodId)
}
