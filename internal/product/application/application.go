package application

import (
	"meliarqsoft2/internal/product/domain/ports"
	application2 "meliarqsoft2/internal/purchase/application"
	"meliarqsoft2/internal/seller/application"
)

type ProductApplication struct {
	repo            ports.IProductRepository
	sellerService   *application.SellerService
	purchaseService *application2.PurchaseService
}

func NewProductApplication(repo ports.IProductRepository, sellerService *application.SellerService,
	purchaseService *application2.PurchaseService) *ProductApplication {
	return &ProductApplication{repo: repo, sellerService: sellerService, purchaseService: purchaseService}
}
