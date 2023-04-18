package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"meliarqsoft2/internal/application/command/action"
	"meliarqsoft2/internal/seller/infrastructure/dto"
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
	var sellerDTO dto.CreateSellerRequest
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
	c.JSON(http.StatusCreated, dto.SellerID{ID: id})
}
