package handler

import (
	"meliarqsoft2/internal/seller/domain/ports"
)

type SellerGinHandler struct {
	service ports.ISellerService
}

func NewSellerGinHandler(service ports.ISellerService) *SellerGinHandler {
	return &SellerGinHandler{service: service}
}
