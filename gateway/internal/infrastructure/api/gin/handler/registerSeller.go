package handler

import (
	"github.com/gin-gonic/gin"
	"meliarqsoft2/internal/domain/model"
	"meliarqsoft2/pkg/exceptions/application"
	"net/http"
)

type GinRegisterSeller struct {
	sellerService model.ISellerService
}

func NewGinRegisterSeller(sellerService model.ISellerService) *GinRegisterSeller {
	return &GinRegisterSeller{sellerService: sellerService}
}

// Execute Register Seller
// @Summary Create a seller
// @Description Create seller
// @Accept json
// @Produce json
// @Tags Sellers
// @Param Body body model.CreateSellerRequest true "Register"
// @Success 201 {object} model.SellerID
// @Failure 400
// @Failure 409
// @Failure 500
// @Failure 503
// @Router /sellers [POST]
func (handler GinRegisterSeller) Execute(c *gin.Context) {
	var sellerReq model.CreateSellerRequest
	if err := c.BindJSON(&sellerReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := handler.sellerService.Register(sellerReq)
	if err != nil {
		application.MeliGinHandlerError{}.Execute(err, c)
		return
	}
	c.JSON(http.StatusCreated, model.SellerID{ID: id})
}
