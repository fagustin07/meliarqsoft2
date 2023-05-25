package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/application/action"
	"meliarqsoft2/internal/infrastructure/api/dto"
	"meliarqsoft2/pkg/exceptions/application"
	"net/http"
)

type GinUpdateCustomer struct {
	UpdateCustomerEvent *action.UpdateCustomerEvent
}

func NewGinUpdateCustomer(unregisterCustomerEvent *action.UpdateCustomerEvent) *GinUpdateCustomer {
	return &GinUpdateCustomer{UpdateCustomerEvent: unregisterCustomerEvent}
}

// Execute Update
// @Summary Update a customer
// @Description Update customer from a customer
// @Produce json
// @Tags Customers
// @Param Body body dto.UpdateCustomerRequest true "Register"
// @Param 	id 	path  string true "ID from customer to update"
// @Success 204
// @Failure 404
// @Failure 400
// @Router /customers/{id} [PUT]
func (handler GinUpdateCustomer) Execute(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var dataToUpdate dto.UpdateCustomerRequest
	if err := c.BindJSON(&dataToUpdate); err != nil {
		application.MeliGinHandlerError{}.Execute(err, c)
		return
	}

	err = handler.UpdateCustomerEvent.Execute(id, dataToUpdate.Name, dataToUpdate.Surname, dataToUpdate.Email)

	if err != nil {
		application.MeliGinHandlerError{}.Execute(err, c)
		return
	}

	c.Status(http.StatusNoContent)
}
