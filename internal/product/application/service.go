package application

import (
	"meliarqsoft2/internal/product/domain/ports"
	"meliarqsoft2/internal/seller/application"
)

type ProductService struct {
	repo          ports.IProductRepository
	sellerService *application.SellerService
}

func NewProductService(repo ports.IProductRepository, sellerService *application.SellerService) *ProductService {
	return &ProductService{repo: repo, sellerService: sellerService}
}
