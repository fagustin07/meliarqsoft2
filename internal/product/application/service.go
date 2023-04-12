package application

import (
	"meliarqsoft2/internal/product/domain/ports"
	application2 "meliarqsoft2/internal/purchase/application"
	"meliarqsoft2/internal/seller/application"
)

type ProductService struct {
	repo            ports.IProductRepository
	sellerService   *application.SellerService
	purchaseService *application2.PurchaseService
}

func NewProductService(repo ports.IProductRepository, sellerService *application.SellerService,
	purchaseService *application2.PurchaseService) *ProductService {
	return &ProductService{repo: repo, sellerService: sellerService, purchaseService: purchaseService}
}
