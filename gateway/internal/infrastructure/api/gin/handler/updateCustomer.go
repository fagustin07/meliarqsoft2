package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/model"
	"meliarqsoft2/pkg/exceptions/application"
	"net/http"
)

type GinUpdateCustomer struct {
	CustomerService model.ICustomerService
}

func NewGinUpdateCustomer(customerService model.ICustomerService) *GinUpdateCustomer {
	return &GinUpdateCustomer{CustomerService: customerService}
}

// Execute Update
// @Summary Update a customer
// @Description Update customer from a customer
// @Produce json
// @Tags Customers
// @Param Body body model.UpdateCustomerRequest true "Update"
// @Param 	id 	path  string true "ID from customer to update"
// @Success 204
// @Failure 400
// @Failure 404
// @Failure 409
// @Failure 500
// @Failure 503
// @Router /customers/{id} [PUT]
func (handler GinUpdateCustomer) Execute(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var dataToUpdate model.UpdateCustomerRequest
	if err := c.BindJSON(&dataToUpdate); err != nil {
		application.MeliGinHandlerError{}.Execute(err, c)
		return
	}

	err = handler.CustomerService.Update(id, dataToUpdate.Name, dataToUpdate.Surname, dataToUpdate.Email)

	if err != nil {
		application.MeliGinHandlerError{}.Execute(err, c)
		return
	}

	c.Status(http.StatusNoContent)
}
