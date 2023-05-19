package handler

import (
	"github.com/gin-gonic/gin"
)

type GinMakePurchase struct {
	//MakePurchaseEvent *action.MakePurchaseEvent
}

//func NewGinMakePurchase(makePurchaseEvent *action.MakePurchaseEvent) *GinMakePurchase {
//	return &GinMakePurchase{MakePurchaseEvent: makePurchaseEvent}
//}

// Execute Create
// @Summary Make a purchase.
// @Description Make User buy Product by ids.
// @Accept json
// @Produce json
// @Tags Purchases
// @Param 	Body body model.CreatePurchaseRequest true "Register"
// @Success 201 {object} []model.Purchase
// @Failure 400
// @Failure 404
// @Failure 406
// @Failure 500
// @Failure 503
// @Router /products/purchases [POST]
func (handler GinMakePurchase) Execute(c *gin.Context) {
	//var purchaseDTO model.CreatePurchaseRequest
	//if err := c.BindJSON(&purchaseDTO); err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}
	//
	//toModel, err := purchaseDTO.MapToModel()
	//if err != nil {
	//	application.MeliGinHandlerError{}.Execute(err, c)
	//	return
	//}
	//id, total, err := handler.MakePurchaseEvent.Execute(&toModel)
	//
	//toModel.ID = id
	//newTotal, err := model.NewTotal(total)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}
	//toModel.Total = newTotal
	//
	//c.JSON(http.StatusCreated, model.MapPurchaseToJSON(&toModel))
}
