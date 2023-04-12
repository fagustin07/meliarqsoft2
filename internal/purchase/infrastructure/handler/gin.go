package handler

import (
	"meliarqsoft2/internal/purchase/domain/ports"
)

type PurchaseGinHandler struct {
	purchaseService ports.IPurchaseService
}

func NewPurchaseGinHandler(manager ports.IPurchaseService) *PurchaseGinHandler {
	return &PurchaseGinHandler{purchaseService: manager}
}
