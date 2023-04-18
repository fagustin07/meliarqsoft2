package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"meliarqsoft2/internal/domain/application/action"
	dto2 "meliarqsoft2/internal/infrastructure/api/dto"
	"net/http"
)

type GinRegisterSeller struct {
	registerSellerEvent *action.RegisterSellerEvent
}

func NewGinRegisterSeller(registerSellerEvent *action.RegisterSellerEvent) *GinRegisterSeller {
	return &GinRegisterSeller{registerSellerEvent: registerSellerEvent}
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

	id, err := handler.registerSellerEvent.Execute(sellerDTO.BusinessName, sellerDTO.Email)
	if err != nil {
		c.Status(http.StatusBadRequest)
		log.Print(err)
		return
	}
	c.JSON(http.StatusCreated, dto2.SellerID{ID: id})
}
