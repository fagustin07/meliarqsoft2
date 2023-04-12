package application

import (
	"meliarqsoft2/internal/product/application"
	"meliarqsoft2/internal/seller/domain/ports"
)

type SellerService struct {
	repo           ports.ISellerRepository
	productService *application.ProductService
}

func NewSellerService(repo ports.ISellerRepository, productService *application.ProductService) *SellerService {
	return &SellerService{repo: repo, productService: productService}
}
