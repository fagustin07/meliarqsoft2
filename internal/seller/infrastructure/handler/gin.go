package handler

import (
	"meliarqsoft2/internal/seller/domain/ports"
)

type SellerGinHandler struct {
	manager ports.ISellerManager
}

func NewSellerGinHandler(manager ports.ISellerManager) *SellerGinHandler {
	return &SellerGinHandler{manager: manager}
}
