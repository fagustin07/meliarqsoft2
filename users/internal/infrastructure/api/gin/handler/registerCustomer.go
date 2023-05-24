package handler

import (
	"github.com/gin-gonic/gin"
	"meliarqsoft2/internal/domain/application/action"
	"meliarqsoft2/internal/domain/model"
	dto2 "meliarqsoft2/internal/infrastructure/api/dto"
	"meliarqsoft2/pkg/exceptions/application"
	"net/http"
)

type GinRegisterCustomer struct {
	RegisterCustomerEvent *action.RegisterCustomerEvent
}

func NewGinCustomerRegister(registerCustomerEvent *action.RegisterCustomerEvent) *GinRegisterCustomer {
	return &GinRegisterCustomer{
		RegisterCustomerEvent: registerCustomerEvent,
	}
}

// Execute Register Customer
// @Summary Create a customer
// @Description Create customer
// @Accept json
// @Produce json
// @Tags Customers
// @Param Customer body dto.CreateCustomerRequest true "Register"
// @Success 201 {object} dto.CustomerID
// @Failure 400
// @Failure 500
// @Router /customers [POST]
func (handler GinRegisterCustomer) Execute(c *gin.Context) {
	var customerDTO dto2.CreateCustomerRequest
	if err := c.BindJSON(&customerDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	toModel, err := customerDTO.MapToModel()
	if err != nil {
		application.MeliGinHandlerError{}.Execute(err, c)
		return
	}

	newNotification, err := model.NewNotification(toModel.Name, toModel.Email.Address)
	if err != nil {
		application.MeliGinHandlerError{}.Execute(err, c)
		return
	}

	id, err := handler.RegisterCustomerEvent.Execute(&toModel, newNotification)
	if err != nil {
		application.MeliGinHandlerError{}.Execute(err, c)
		return
	}

	c.JSON(http.StatusCreated, dto2.CustomerID{ID: id})
}
