package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/application/action"
	"meliarqsoft2/pkg/exceptions/application"
	"net/http"
)

type GinRestoreProducts struct {
	restoreProductsEvent *action.RestoreProductsEvent
}

func NewGinRestoreProducts(restoreProductsEvent *action.RestoreProductsEvent) *GinRestoreProducts {
	return &GinRestoreProducts{restoreProductsEvent: restoreProductsEvent}
}

// Execute Restore
// @Summary Restore deleted products
// @Description Restore products
// @Accept json
// @Produce json
// @Tags Products
// @Param IDs body ProductsIDs true "IDs from products to restore"
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /products/restore [POST]
func (handler GinRestoreProducts) Execute(c *gin.Context) {
	var idsToRestore ProductsIDs
	if err := c.BindJSON(&idsToRestore); err != nil {
		application.MeliGinHandlerError{}.Execute(err, c)
		return
	}

	restored, err := handler.restoreProductsEvent.Execute(idsToRestore.IDs)
	if err != nil {
		application.MeliGinHandlerError{}.Execute(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"restored": restored})
}

type ProductsIDs struct {
	IDs []uuid.UUID `json:"ids" bson:"ids" binding:"required"`
}
