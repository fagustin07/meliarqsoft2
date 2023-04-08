package handler

import (
	"meliarqsoft2/internal/purchase/domain/ports"
)

type PurchaseGinHandler struct {
	manager ports.IPurchaseManager
}

func NewPurchaseGinHandler(manager ports.IPurchaseManager) *PurchaseGinHandler {
	return &PurchaseGinHandler{manager: manager}
}
