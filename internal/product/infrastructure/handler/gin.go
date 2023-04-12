package handler

import (
	"meliarqsoft2/internal/product/domain/ports"
)

type ProductGinHandler struct {
	productService ports.IProductService
}

func NewProductGinHandler(service ports.IProductService) *ProductGinHandler {
	return &ProductGinHandler{productService: service}
}
