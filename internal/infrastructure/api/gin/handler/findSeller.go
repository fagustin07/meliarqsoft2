package handler

import (
	"github.com/gin-gonic/gin"
	"meliarqsoft2/internal/domain/application/query"
	"meliarqsoft2/internal/infrastructure/api/dto"
	"meliarqsoft2/pkg/exceptions/application"
	"net/http"
)

type GinFindSeller struct {
	FindSellerEvent *query.FindSellerEvent
}

func NewGinFindSeller(findSellerEvent *query.FindSellerEvent) *GinFindSeller {
	return &GinFindSeller{FindSellerEvent: findSellerEvent}
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

	resp, err := handler.FindSellerEvent.Execute(name)
	if err != nil {
		application.MeliGinHandlerError{}.Execute(err, c)
		return
	}

	var res []dto.SellerDTO
	for _, elem := range resp {
		res = append(res, dto.MapSellerToJSON(elem))
	}
	c.JSON(http.StatusOK, res)
}
