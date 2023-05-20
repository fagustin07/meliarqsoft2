package handler

import (
	"github.com/gin-gonic/gin"
)

type GinCreateProduct struct {
	//CreateProductEvent *action.CreateProductEvent
}

//func NewGinCreateProduct(createProductEvent *action.CreateProductEvent) *GinCreateProduct {
//	return &GinCreateProduct{CreateProductEvent: createProductEvent}
//}

// Execute Create
// @Summary Create a product
// @Description Create product from a seller
// @Accept json
// @Produce json
// @Tags Products
// @Param Body body model.CreateProductRequest true "Create a product"
// @Success 200 {object} model.ProductID
// @Failure 400
// @Failure 404
// @Failure 409
// @Failure 500
// @Failure 503
// @Router /products [post]
func (handler GinCreateProduct) Execute(c *gin.Context) {
	//var productDTO dto2.CreateProductRequest
	//if err := c.BindJSON(&productDTO); err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}
	//
	//product, err := productDTO.MapToModel()
	//if err != nil {
	//	application.MeliGinHandlerError{}.Execute(err, c)
	//	return
	//}

	//createdProductID, err := handler.CreateProductEvent.Execute(product)
	//if err != nil {
	//	application.MeliGinHandlerError{}.Execute(err, c)
	//	return
	//}
	//c.JSON(http.StatusCreated, dto2.ProductID{ID: createdProductID})
}
