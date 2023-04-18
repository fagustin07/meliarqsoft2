package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"meliarqsoft2/internal/domain/application/query"
	"meliarqsoft2/internal/infrastructure/api/dto"
	"net/http"
)

type GinFindPurchases struct {
	FindPurchasesFromProductEvent *query.FindPurchasesFromProductEvent
}

func NewGinFindPurchases(findPurchasesFromProductEvent *query.FindPurchasesFromProductEvent) *GinFindPurchases {
	return &GinFindPurchases{FindPurchasesFromProductEvent: findPurchasesFromProductEvent}
}

// Execute Find
// @Summary Find purchases from product.
// @Description Find purchases from product
// @Produce json
// @Tags Purchases
// @Param 	id 	path  string true "ID from product"
// @Success 200
// @Failure 404
// @Failure 400
// @Router /products/{id}/purchases [GET]
func (handler GinFindPurchases) Execute(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	log.Println(id)
	if err != nil {
		c.Status(http.StatusBadRequest)
		log.Println(err.Error())
		return
	}
	resp, err := handler.FindPurchasesFromProductEvent.Execute(id)
	if err != nil {
		c.Status(http.StatusBadRequest)
		log.Print(err)
		return
	}

	var res []dto.PurchaseDTO
	for _, elem := range resp {
		res = append(res, dto.MapPurchaseToJSON(elem))
	}

	c.JSON(http.StatusOK, res)
}
