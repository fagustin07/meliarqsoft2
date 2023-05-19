package handler

import (
	"github.com/gin-gonic/gin"
	"meliarqsoft2/internal/domain/application/query"
	"meliarqsoft2/internal/infrastructure/api/dto"
	"meliarqsoft2/pkg/exceptions/application"
	"net/http"
)

type GinFindCustomer struct {
	FindCustomerEvent *query.FindCustomerEvent
}

func NewGinFindCustomer(findCustomerEvent *query.FindCustomerEvent) *GinFindCustomer {
	return &GinFindCustomer{FindCustomerEvent: findCustomerEvent}
}

// Execute Find Customer
// @Summary Find customer
// @Description Find customer with given id
// @Accept json
// @Produce json
// @Tags Customers
// @Param email	query string  false "find name"
// @Success 200 {object} []dto.CustomerDTO
// @Failure 400
// @Failure 500
// @Router /customers [GET]
func (handler GinFindCustomer) Execute(c *gin.Context) {
	emailPattern := c.Query("email")

	resp, err := handler.FindCustomerEvent.Execute(emailPattern)
	if err != nil {
		application.MeliGinHandlerError{}.Execute(err, c)
		return
	}

	var res []dto.CustomerDTO
	for _, elem := range resp {
		res = append(res, dto.MapCustomerToJSON(elem))
	}

	c.JSON(http.StatusOK, res)
}
