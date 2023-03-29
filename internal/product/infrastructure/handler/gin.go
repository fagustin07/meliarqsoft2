package handler

import (
	"meliarqsoft2/internal/product/domain/ports"
)

type ProductGinHandler struct {
	app ports.IProductApplication
}

func NewProductGinHandler(app ports.IProductApplication) *ProductGinHandler {
	return &ProductGinHandler{app: app}
}
