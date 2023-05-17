package handler

import (
	"github.com/gin-gonic/gin"
	"meliarqsoft2/internal/domain/model"
	"meliarqsoft2/pkg/exceptions/application"
	"net/http"
)

type GinFindSeller struct {
	sellerService model.ISellerService
}

func NewGinFindSeller(sellerService model.ISellerService) *GinFindSeller {
	return &GinFindSeller{sellerService: sellerService}
}

// Execute Find
// @Summary Find sellers
// @Description Find sellers that contains given business name
// @Accept json
// @Produce json
// @Tags Sellers
// @Param business_name	query string  false "find name"
// @Success 200 {object} []dto.SellerDTO
// @Failure 400
// @Failure 500
// @Router /sellers [get]
func (handler GinFindSeller) Execute(c *gin.Context) {
	name := c.Query("business_name")

	resp, err := handler.sellerService.Find(name)
	if err != nil {
		application.MeliGinHandlerError{}.Execute(err, c)
		return
	}

	c.JSON(http.StatusOK, resp)
}
