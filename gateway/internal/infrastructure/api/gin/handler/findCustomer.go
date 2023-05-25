package handler

import (
	"github.com/gin-gonic/gin"
	"meliarqsoft2/internal/domain/model"
	"meliarqsoft2/pkg/exceptions/application"
	"net/http"
)

type GinFindCustomer struct {
	CustomerService model.ICustomerService
}

func NewGinFindCustomer(customerService model.ICustomerService) *GinFindCustomer {
	return &GinFindCustomer{CustomerService: customerService}
}

// Execute Find Customer
// @Summary Find customer
// @Description Find customer with given id
// @Accept json
// @Produce json
// @Tags Customers
// @Param email	query string  false "find name"
// @Success 200 {object} []model.Customer
// @Failure 400
// @Failure 500
// @Router /customers [GET]
func (handler GinFindCustomer) Execute(c *gin.Context) {
	emailPattern := c.Query("email")

	resp, err := handler.CustomerService.Find(emailPattern)
	if err != nil {
		application.MeliGinHandlerError{}.Execute(err, c)
		return
	}

	c.JSON(http.StatusOK, resp)
}
