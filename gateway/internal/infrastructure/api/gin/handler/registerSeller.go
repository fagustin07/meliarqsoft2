package handler

import (
	"github.com/gin-gonic/gin"
	"meliarqsoft2/internal/domain/model"
	dto2 "meliarqsoft2/internal/infrastructure/api/dto"
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
// @Param Body body dto.CreateSellerRequest true "Register"
// @Success 201 {object} dto.SellerID
// @Failure 400
// @Failure 500
// @Router /sellers [POST]
func (handler GinRegisterSeller) Execute(c *gin.Context) {
	var sellerDTO dto2.CreateSellerRequest
	if err := c.BindJSON(&sellerDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	requestedSeller, err := sellerDTO.MapToModel()
	if err != nil {
		application.MeliGinHandlerError{}.Execute(err, c)
		return
	}

	id, err := handler.sellerService.Register(&requestedSeller)
	if err != nil {
		application.MeliGinHandlerError{}.Execute(err, c)
		return
	}
	c.JSON(http.StatusCreated, dto2.SellerID{ID: id})
}
