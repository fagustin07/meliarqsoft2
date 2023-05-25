package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/application/action"
	"meliarqsoft2/pkg/exceptions/application"
	"net/http"
)

type GinUnregisterCustomer struct {
	UnregisterCustomerEvent *action.UnregisterCustomerEvent
}

func NewGinUnregisterCustomer(unregisterCustomerEvent *action.UnregisterCustomerEvent) *GinUnregisterCustomer {
	return &GinUnregisterCustomer{UnregisterCustomerEvent: unregisterCustomerEvent}
}

// Execute Unregister Customer
// @Summary Unregister a customer
// @Description Unregister customer from a customer
// @Produce json
// @Tags Customers
// @Param 	id 	path  string true "ID from customer to unregister"
// @Success 204
// @Failure 400
// @Failure 404
// @Failure 400
// @Router /customers/{id} [DELETE]
func (handler GinUnregisterCustomer) Execute(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = handler.UnregisterCustomerEvent.Execute(id)

	if err != nil {
		application.MeliGinHandlerError{}.Execute(err, c)
		return
	}

	c.Status(http.StatusNoContent)
}
