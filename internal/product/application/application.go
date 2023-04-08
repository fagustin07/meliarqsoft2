package application

import (
	"meliarqsoft2/internal/product/domain/ports"
	"meliarqsoft2/internal/seller/application"
)

type ProductApplication struct {
	repo          ports.IProductRepository
	sellerManager *application.SellerManager
}

func NewProductApplication(repo ports.IProductRepository, manager *application.SellerManager) *ProductApplication {
	return &ProductApplication{repo: repo, sellerManager: manager}
}
