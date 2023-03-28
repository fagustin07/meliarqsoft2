package handler

import (
	"github.com/gin-gonic/gin"
	"meliarqsoft2/internal/product/domain/ports"
)

type ProductGinHandler struct {
	app ports.IProductApplication
}

func NewProductGinHandler(app ports.IProductApplication) *ProductGinHandler {
	return &ProductGinHandler{app: app}
}

func (p ProductGinHandler) Update(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p ProductGinHandler) Delete(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
