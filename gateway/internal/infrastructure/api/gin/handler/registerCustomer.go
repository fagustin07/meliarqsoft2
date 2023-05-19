package handler

import (
	"github.com/gin-gonic/gin"
	"meliarqsoft2/internal/domain/model"
	"meliarqsoft2/pkg/exceptions/application"
	"net/http"
)

type GinRegisterCustomer struct {
	CustomerService model.ICustomerService
}

func NewGinCustomerRegister(customerService model.ICustomerService) *GinRegisterCustomer {
	return &GinRegisterCustomer{CustomerService: customerService}
}

// Execute Register Customer
// @Summary Create a customer
// @Description Create customer
// @Accept json
// @Produce json
// @Tags Customers
// @Param Body body model.CreateCustomerRequest true "Register"
// @Success 201 {object} model.CustomerID
// @Failure 400
// @Failure 409
// @Failure 500
// @Failure 503
// @Router /customers [POST]
func (handler GinRegisterCustomer) Execute(c *gin.Context) {
	var customerReq model.CreateCustomerRequest
	if err := c.BindJSON(&customerReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := handler.CustomerService.Register(customerReq)
	if err != nil {
		application.MeliGinHandlerError{}.Execute(err, c)
		return
	}
	c.JSON(http.StatusCreated, model.CustomerID{ID: id})
}
