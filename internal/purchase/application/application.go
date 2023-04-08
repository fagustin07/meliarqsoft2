package application

import (
	"meliarqsoft2/internal/product/application"
	"meliarqsoft2/internal/purchase/domain/ports"
	application2 "meliarqsoft2/internal/user/application"
)

type PurchaseManager struct {
	repository     ports.IPurchaseRepository
	productManager *application.ProductApplication
	userManager    *application2.UserManager
}

func NewPurchaseManager(repo ports.IPurchaseRepository, productManager *application.ProductApplication,
	userManager *application2.UserManager) *PurchaseManager {
	return &PurchaseManager{repository: repo, productManager: productManager, userManager: userManager}
}
