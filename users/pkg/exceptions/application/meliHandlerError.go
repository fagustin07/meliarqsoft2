package application

import (
	"github.com/gin-gonic/gin"
	"meliarqsoft2/pkg/exceptions/model"
	"net/http"
)

type MeliGinHandlerError struct{}

func (handlerErr MeliGinHandlerError) Execute(err error, c *gin.Context) {
	if isNotFoundError(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if isAlreadyExistError(err) {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	if isBusinessError(err) || isNotAcceptableError(err) {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}

func isNotFoundError(err error) bool {
	_, customer := err.(model.CustomerNotFoundError)
	_, purchase := err.(model.PurchaseNotFoundError)
	_, seller := err.(model.SellerNotFoundError)
	_, product := err.(model.ProductNotFound)

	return customer || purchase || seller || product
}

func isAlreadyExistError(err error) bool {
	_, customer := err.(model.CustomerAlreadyExistError)
	_, seller := err.(model.SellerAlreadyExist)
	_, product := err.(model.ProductAlreadyExist)

	return customer || seller || product
}

func isBusinessError(err error) bool {
	_, email := err.(model.InvalidEmailError)
	_, price := err.(model.InvalidPriceError)
	_, stock := err.(model.InvalidStockError)
	_, units := err.(model.InvalidUnitsError)
	_, total := err.(model.InvalidTotalPriceError)

	return email || price || stock || units || total
}

func isNotAcceptableError(err error) bool {
	_, filterPrice := err.(model.MinAndMaxPriceCombinationError)

	return filterPrice
}
