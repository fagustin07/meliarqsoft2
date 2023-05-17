package application

import (
	"github.com/gin-gonic/gin"
	"log"
	"meliarqsoft2/pkg/exceptions/model"
	"net/http"
)

type MeliGinHandlerError struct{}

func (handlerErr MeliGinHandlerError) Execute(err error, c *gin.Context) {
	log.Println(err.Error())
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
	_, user := err.(model.UserNotFoundError)
	_, purchase := err.(model.PurchaseNotFoundError)
	_, seller := err.(model.SellerNotFoundError)
	_, product := err.(model.ProductNotFound)

	return user || purchase || seller || product
}

func isAlreadyExistError(err error) bool {
	_, user := err.(model.UserAlreadyExistError)
	_, seller := err.(model.SellerAlreadyExist)
	_, product := err.(model.ProductAlreadyExist)

	return user || seller || product
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
