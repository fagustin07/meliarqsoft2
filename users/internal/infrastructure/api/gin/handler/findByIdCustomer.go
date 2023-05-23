package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/application/query"
	"meliarqsoft2/internal/infrastructure/api/dto"
	"meliarqsoft2/pkg/exceptions/application"
	"net/http"
)

type GinFindByIdCustomer struct {
	FindByIdCustomerEvent *query.FindCustomerByIdEvent
}

func NewGinFindByIdCustomer(findByIdCustomerEvent *query.FindCustomerByIdEvent) *GinFindByIdCustomer {
	return &GinFindByIdCustomer{FindByIdCustomerEvent: findByIdCustomerEvent}
}

// Execute FindById customer
// @Summary FindById a customer
// @Description FindById customer
// @Accept json
// @Produce json
// @Tags Customers
// @Param 	id 	path  string true "ID from customer"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /customers/{id} [GET]
func (handler GinFindByIdCustomer) Execute(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer, err := handler.FindByIdCustomerEvent.Execute(id)

	if err != nil {
		application.MeliGinHandlerError{}.Execute(err, c)
		return
	}

	c.JSON(http.StatusOK, dto.MapCustomerToJSON(customer))
}
